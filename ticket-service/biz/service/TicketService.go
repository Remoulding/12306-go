package service

import (
	"context"
	"fmt"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/Remoulding/12306-go/ticket-service/tools"
	"github.com/bytedance/sonic"
	"github.com/go-redsync/redsync/v4"
	"github.com/samber/lo"
	"slices"
	"strconv"
	"strings"
	"time"
	//"github.com/Remoulding/12306-go/ticket-service/biz/dao"
)

type TicketService struct {
}

func NewTicketService() *TicketService {
	return &TicketService{}
}

func (s *TicketService) PageListTicketQuery(ctx context.Context, req *ticket_service.TicketPageQueryRequest) (*ticket_service.TicketPageQueryResponse, error) {
	resp := &ticket_service.TicketPageQueryResponse{}
	stationDetailMap, err := GetHashCache(ctx, configs.RegionTrainStationMapping)
	if err != nil || stationDetailMap == nil {
		return resp, nil
	}
	startRegion := stationDetailMap[req.FromStation]
	endRegion := stationDetailMap[req.ToStation]
	if startRegion == "" || endRegion == "" {
		stations, err := dao.QueryStations(ctx, nil)
		if err != nil {
			return resp, nil
		}
		for _, item := range stations {
			stationDetailMap[item.Code] = item.RegionName
		}
		err = SetHashCache(ctx, configs.RegionTrainStationMapping, stationDetailMap)
		if err != nil {
			return resp, nil
		}
		startRegion = stationDetailMap[req.FromStation]
		endRegion = stationDetailMap[req.ToStation]
	}

	regionStationHashKey := fmt.Sprintf(configs.RegionTrainStation, startRegion, endRegion)
	loadData, err := SafeLoad(ctx, regionStationHashKey, configs.LockRegionTrainStation, func(ctx context.Context) (interface{}, error) {
		trainInfoMap := make(map[string]string)
		conditions := map[string]interface{}{
			"start_region = ?": startRegion,
			"end_region = ?":   endRegion,
		}
		trainStationRelations, err := dao.QueryTrainStationRelation(ctx, conditions)
		if err != nil {
			return trainInfoMap, err
		}
		for _, item := range trainStationRelations {
			trainID := strconv.Itoa(int(item.TrainID))
			trainData, err := SafeLoad(ctx, configs.TrainInfo+trainID, configs.LockTrain+trainID, func(ctx context.Context) (interface{}, error) {
				return dao.QueryTrainById(ctx, item.TrainID)
			})
			if err != nil {
				log.WithContext(ctx).Errorf("query train info failed, err: %v", err)
				return nil, err
			}
			var train *model.TrainDO
			_ = sonic.UnmarshalString(trainData, &train)
			trainInfo := &ticket_service.TicketInfo{}
			trainInfo.TrainId = trainID
			trainInfo.TrainNumber = train.TrainNumber
			trainInfo.DepartureTime = tools.ConvertTimeToStr(train.DepartureTime)
			trainInfo.ArrivalTime = tools.ConvertTimeToStr(train.ArrivalTime)
			trainInfo.Duration = tools.CalculateHourDifference(train.DepartureTime, train.ArrivalTime)
			trainInfo.Departure = item.Departure
			trainInfo.Arrival = item.Arrival
			trainInfo.DepartureFlag = item.DepartureFlag
			trainInfo.ArrivalFlag = item.ArrivalFlag
			trainInfo.TrainType = int32(train.TrainType)
			trainInfo.TrainBrand = train.TrainBrand
			if len(train.TrainTag) > 0 {
				trainInfo.TrainTags = strings.Split(train.TrainTag, ",")
			}
			if time.Now().Before(train.DepartureTime) {
				trainInfo.SaleStatus = 1 // 未开售
			}
			trainInfo.DaysArrived = int32((train.ArrivalTime.Hour() - train.DepartureTime.Hour()) / 24)
			trainInfo.SaleTime = train.ArrivalTime.Format("01-02 15:04")
			tranInfoKey := fmt.Sprintf("%d_%s_%s", item.TrainID, item.Departure, item.Arrival)
			trainInfoData, _ := sonic.MarshalString(trainInfo)
			trainInfoMap[tranInfoKey] = trainInfoData
		}
		return trainInfoMap, nil
	})
	if err != nil {
		return resp, nil
	}
	var trainInfoMap map[string]string
	_ = sonic.UnmarshalString(loadData, &trainInfoMap)
	trainList := make([]*ticket_service.TicketInfo, 0, len(trainInfoMap))
	for _, item := range trainInfoMap {
		trainInfo := &ticket_service.TicketInfo{}
		_ = sonic.UnmarshalString(item, trainInfo)
		trainList = append(trainList, trainInfo)
	}
	slices.SortStableFunc(trainList, func(a, b *ticket_service.TicketInfo) int {
		return strings.Compare(a.DepartureTime, b.DepartureTime)
	})

	resp.Data = s.buildPageListTicketQueryData(trainList)
	resp.Success = true
	return resp, nil
}

func (s *TicketService) buildPageListTicketQueryData(trainList []*ticket_service.TicketInfo) *ticket_service.TicketPageQueryData {
	data := &ticket_service.TicketPageQueryData{}
	data.TrainList = trainList
	for _, item := range trainList {
		data.DepartureStationList = append(data.DepartureStationList, item.GetDeparture())
		data.ArrivalStationList = append(data.ArrivalStationList, item.GetArrival())
		for _, seatClassInfo := range item.GetSeatClassList() {
			data.SeatClassTypeList = append(data.SeatClassTypeList, seatClassInfo.GetType())
		}
	}
	buildBrandList := func(trainList []*ticket_service.TicketInfo) []int32 {
		brandMap := map[int32]struct{}{}
		for _, item := range trainList {
			ss := strings.Split(item.GetTrainBrand(), ",")
			for _, brandStr := range ss {
				brand, _ := strconv.ParseInt(brandStr, 10, 32)
				brandMap[int32(brand)] = struct{}{}
			}
		}
		return lo.MapToSlice(brandMap, func(key int32, val struct{}) int32 {
			return key
		})
	}
	data.TrainBrandList = buildBrandList(trainList)
	return data
}

func (s *TicketService) PurchaseTickets(ctx context.Context, req *ticket_service.PurchaseTicketRequest) (*ticket_service.PurchaseTicketResponse, error) {
	lockKey := fmt.Sprintf(configs.LockPurchaseTickets, req.GetTrainId())
	mutex := rs.NewMutex(lockKey, redsync.WithExpiry(5*time.Second))
	_ = mutex.Lock()
	defer func(mutex *redsync.Mutex) {
		if _, err := mutex.Unlock(); err != nil {
			log.WithContext(ctx).Errorf("SafeGet failed, err: %v", err)
		}
	}(mutex)
	resp := &ticket_service.PurchaseTicketResponse{}
	//data := &ticket_service.TicketPurchaseData{}
	trainData, err := SafeLoad(ctx, configs.TrainInfo+req.GetTrainId(), configs.LockTrain+req.GetTrainId(), func(ctx context.Context) (interface{}, error) {
		trainID, _ := strconv.ParseInt(req.GetTrainId(), 10, 64)
		return dao.QueryTrainById(ctx, trainID)
	})
	if err != nil {
		return resp, nil
	}
	var train model.TrainDO
	_ = sonic.UnmarshalString(trainData, &train)

	return resp, nil
}

func (s *TicketService) SelectSeat(ctx context.Context, trainType int, req *ticket_service.PurchaseTicketRequest) []*TrainPurchaseTicketDTO {
	passengers := req.GetPassengers()
	seatTypeMap := make(map[int32][]*ticket_service.PurchaseTicketPassengerInfo)
	for _, passenger := range passengers {
		seatTypeMap[passenger.SeatType] = append(seatTypeMap[passenger.SeatType], passenger)
	}
	for seatType, list := range seatTypeMap {
		switch seatType {
		case 0:
			s.selectBusinessSeat(ctx, list, req)
		case 1:
			s.selectFirstSeat(ctx, list, req)
		case 2:
			s.selectSecondSeat(ctx, list, req)
		}
	}
	return nil
}
func (s *TicketService) selectBusinessSeat(ctx context.Context, list []*ticket_service.PurchaseTicketPassengerInfo, req *ticket_service.PurchaseTicketRequest) []*TrainPurchaseTicketDTO {
	return nil
}

func (s *TicketService) selectFirstSeat(ctx context.Context, list []*ticket_service.PurchaseTicketPassengerInfo, req *ticket_service.PurchaseTicketRequest) []*TrainPurchaseTicketDTO {
	return nil
}
func (s *TicketService) selectSecondSeat(ctx context.Context, list []*ticket_service.PurchaseTicketPassengerInfo, req *ticket_service.PurchaseTicketRequest) []*TrainPurchaseTicketDTO {
	return nil
}

func (s *TicketService) CancelTicketOrder(ctx context.Context, req *ticket_service.CancelTicketOrderRequest) (*ticket_service.CancelTicketOrderResponse, error) {
	//TODO implement me
	// 回滚一下就完事
	panic("implement me")
}
