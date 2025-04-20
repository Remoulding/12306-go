package service

import (
	"context"
	"fmt"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/idl-gen/user_service"
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
	"sync"
	"time"
	//"github.com/Remoulding/12306-go/ticket-service/biz/dao"
)

type TicketService struct {
	seatService *SeatService
}

func NewTicketService() *TicketService {
	return &TicketService{
		seatService: NewSeatService(),
	}
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
	// 拿到出发站和到达站的区域
	regionStationHashKey := fmt.Sprintf(configs.RegionTrainStation, startRegion, endRegion, req.DepartureDate)
	lockKey := fmt.Sprintf(configs.LockRegionTrainStation, startRegion, endRegion, req.DepartureDate)
	loadData, err := SafeLoad(ctx, regionStationHashKey, lockKey, func(ctx context.Context) (interface{}, error) {
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
			trainInfo.DepartureTime, trainInfo.ArrivalTime = tools.ComputeFullTimes(req.GetDepartureDate(), train.DepartureTime, train.ArrivalTime, train.CrossDay)
			trainInfo.Duration = tools.CalculateHourDifference(trainInfo.DepartureTime, trainInfo.ArrivalTime)
			trainInfo.Departure = item.Departure
			trainInfo.Arrival = item.Arrival
			trainInfo.DepartureFlag = item.DepartureFlag
			trainInfo.ArrivalFlag = item.ArrivalFlag
			trainInfo.DaysArrived = int32(train.CrossDay)
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
	// 获取座位信息
	wg := sync.WaitGroup{}
	var err2 error
	for _, item := range trainList {
		wg.Add(1)
		go func() {
			defer wg.Done()
			redisKey := fmt.Sprintf(configs.TrainStationPrice, item.TrainId, item.Departure, item.Arrival)
			loadData, err = SafeLoad(ctx, redisKey, "", func(ctx context.Context) (interface{}, error) {
				prices, err := dao.QueryTrainStationPrice(ctx, map[string]interface{}{
					"train_id":  item.TrainId,
					"departure": item.Departure,
					"arrival":   item.Arrival,
				})
				if err != nil {
					return nil, err
				}
				mp := make(map[int]int)
				for i := range prices {
					mp[prices[i].SeatType] = prices[i].Price
				}
				return mp, nil
			})
			if err != nil {
				log.WithContext(ctx).Errorf("query train station price failed, err: %v", err)
				err2 = err
				return
			}
			seatMap := make(map[int]int)
			_ = sonic.UnmarshalString(loadData, &seatMap)
			redisKey = fmt.Sprintf(configs.TrainRemainTicket, item.TrainId, item.Departure, item.Arrival, req.GetDepartureDate())
			lockKey = fmt.Sprintf(configs.LockTrainRemainTicket, item.TrainId, item.Departure, item.Arrival, req.GetDepartureDate())
			trainID, _ := strconv.ParseInt(item.TrainId, 10, 64)
			loadHashData, err := SafeLoadHash(ctx, redisKey, lockKey, func(ctx context.Context) (map[string]string, error) {
				groupedByType, err3 := dao.QuerySeatCountGroupedByType(ctx, trainID, item.Departure, item.Arrival, req.GetDepartureDate())
				if err != nil {
					return nil, err3
				}
				mp := make(map[string]string)
				for _, seatCount := range groupedByType {
					mp[strconv.Itoa(seatCount.SeatType)] = strconv.Itoa(int(seatCount.Count))
				}
				return mp, nil
			})
			if err != nil {
				log.WithContext(ctx).Errorf("query seat count failed, err: %v", err)
				err2 = err
				return
			}
			seatClassInfoList := make([]*ticket_service.SeatClassInfo, 0, len(seatMap))
			for seatType, price := range seatMap {
				if seatType > 2 {
					continue
				}
				quantity, _ := strconv.ParseInt(loadHashData[strconv.Itoa(seatType)], 10, 64)
				seatClassInfo := &ticket_service.SeatClassInfo{
					Type:      int32(seatType),
					Price:     float64(price) / 100,
					Candidate: false,
					Quantity:  int32(quantity),
				}
				// 根据 type 从小到大排序
				slices.SortStableFunc(seatClassInfoList, func(a, b *ticket_service.SeatClassInfo) int {
					return int(a.Type - b.Type)
				})
				seatClassInfoList = append(seatClassInfoList, seatClassInfo)
			}
			item.SeatClassList = seatClassInfoList
		}()
	}
	wg.Wait()
	if err2 != nil {
		log.WithContext(ctx).Errorf("query train station seat failed, err: %v", err2)
		resp.Message = "查询失败"
		return resp, nil
	}
	resp.Data = s.buildPageListTicketQueryData(trainList)
	resp.Success = true
	return resp, nil
}

func (s *TicketService) buildPageListTicketQueryData(trainList []*ticket_service.TicketInfo) *ticket_service.TicketPageQueryData {
	data := &ticket_service.TicketPageQueryData{}
	data.TrainList = trainList
	var departureStationList, arrivalStationList []string
	var seatClassTypeList []int32
	for _, item := range trainList {
		departureStationList = append(departureStationList, item.GetDeparture())
		arrivalStationList = append(arrivalStationList, item.GetArrival())
		for _, seatClassInfo := range item.GetSeatClassList() {
			seatClassTypeList = append(seatClassTypeList, seatClassInfo.GetType())
		}
	}
	data.DepartureStationList = lo.Uniq(departureStationList)
	data.ArrivalStationList = lo.Uniq(arrivalStationList)
	data.SeatClassTypeList = lo.Uniq(seatClassTypeList)

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
	resp := &ticket_service.PurchaseTicketResponse{}
	lockKey := fmt.Sprintf(configs.LockPurchaseTickets, req.GetTrainId())
	mutex := configs.Rs.NewMutex(lockKey, redsync.WithExpiry(5*time.Second))
	if err := mutex.Lock(); err != nil {
		log.WithContext(ctx).Errorf("SafeGet add mutex failed, err: %v", err)
		resp.Message = "预占座位失败"
		return resp, nil
	}
	defer func(mutex *redsync.Mutex) {
		if _, err := mutex.Unlock(); err != nil {
			log.WithContext(ctx).Errorf("SafeGet failed, err: %v", err)
		}
	}(mutex)
	//data := &ticket_service.TicketPurchaseData{}
	trainData, err := SafeLoad(ctx, configs.TrainInfo+req.GetTrainId(), configs.LockTrain+req.GetTrainId(), func(ctx context.Context) (interface{}, error) {
		trainID, _ := strconv.ParseInt(req.GetTrainId(), 10, 64)
		return dao.QueryTrainById(ctx, trainID)
	})
	if err != nil {
		log.WithContext(ctx).Errorf("query train info failed, err: %v", err)
		resp.Message = "预占座位失败"
		return resp, nil
	}

	var train model.TrainDO
	if err = sonic.UnmarshalString(trainData, &train); err != nil {
		log.WithContext(ctx).Errorf("unmarshal train info failed, err: %v", err)
		resp.Message = "预占座位失败"
		return resp, nil
	}
	seats, err := s.selectSeat(ctx, req)
	if err != nil {
		log.WithContext(ctx).Errorf("select seat failed, err: %v", err)
		resp.Message = "预占座位失败"
		return resp, nil
	}
	// 生成 Ticket
	ticketList := make([]*model.TicketDO, 0, len(seats))
	for _, seat := range seats {
		depTime, arrTime := tools.ComputeFullTimes(req.GetDepartureDate(), train.DepartureTime, train.ArrivalTime, train.CrossDay)
		pid, _ := strconv.ParseInt(seat.PassengerId, 10, 64)
		ticketList = append(ticketList, &model.TicketDO{
			TrainID:        train.ID,
			TrainNumber:    train.TrainNumber,
			CarriageNumber: seat.CarriageNumber,
			SeatNumber:     seat.SeatNumber,
			Departure:      req.GetDeparture(),
			Arrival:        req.GetArrival(),
			DepartureTime:  tools.ConvertStrToTime(depTime, true),
			ArrivalTime:    tools.ConvertStrToTime(arrTime, true),
			PassengerID:    pid,
			TicketStatus:   model.TicketValid,
			Username:       req.GetUsername(),
		})
	}
	if err = dao.BatchInsertTicket(ctx, ticketList); err != nil {
		log.WithContext(ctx).Errorf("batch insert ticket failed, err: %v", err)
		resp.Message = "预占座位失败"
		return resp, nil
	}
	//tx.Commit()
	data := make([]*ticket_service.TicketOrderDetail, 0, len(seats))
	for _, seat := range seats {
		data = append(data, &ticket_service.TicketOrderDetail{
			SeatType:       int32(seat.SeatType),
			SeatNumber:     seat.SeatNumber,
			CarriageNumber: seat.CarriageNumber,
			Amount:         int32(seat.Amount),
			RealName:       seat.RealName,
			IdCard:         seat.IdCard,
		})
	}
	resp.Data = &ticket_service.TicketPurchaseData{
		TicketOrderDetails: data,
	}
	resp.Success = true
	return resp, nil
}

func (s *TicketService) selectSeat(ctx context.Context, req *ticket_service.PurchaseTicketRequest) ([]*TrainPurchaseTicketDTO, error) {
	passengers := req.GetPassengers()
	passengerIds := make([]int64, 0, len(passengers))
	seatTypeMap := make(map[int32][]*ticket_service.PurchaseTicketPassengerInfo)
	for _, passenger := range passengers {
		pid, _ := strconv.ParseInt(passenger.GetPassengerId(), 10, 64)
		passengerIds = append(passengerIds, pid)
		seatTypeMap[passenger.SeatType] = append(seatTypeMap[passenger.SeatType], passenger)
	}
	resp := make([]*TrainPurchaseTicketDTO, 0, len(passengers))
	for seatType, passengerList := range seatTypeMap {
		seats, err := NewSeatSelectorService(ctx, int(seatType)).SelectSeats(passengerList, req)
		if err != nil {
			log.WithContext(ctx).Errorf("NewSeatSelectorService.select seat failed, err: %v", err)
			return nil, err
		}
		resp = append(resp, seats...)
	}
	log.WithContext(ctx).Infof("NewSeatSelectorService.select seats: %+v", tools.MustJson(resp))
	lockKey := fmt.Sprintf(configs.LockTrainStationPrice, req.GetTrainId(), req.GetDeparture(), req.GetArrival())
	loadData, err := SafeLoad(ctx, fmt.Sprintf(configs.TrainStationPrice, req.GetTrainId(), req.GetDeparture(), req.GetArrival()), lockKey, func(ctx context.Context) (interface{}, error) {
		trainStationPrices, err := dao.QueryTrainStationPrice(ctx, map[string]interface{}{
			"train_id":  req.GetTrainId(),
			"departure": req.GetDeparture(),
			"arrival":   req.GetArrival(),
		})
		if err != nil {
			return nil, err
		}
		mp := make(map[int]int)
		for i := range trainStationPrices {
			mp[trainStationPrices[i].SeatType] = trainStationPrices[i].Price
		}
		return mp, nil
	})
	if err != nil {
		log.WithContext(ctx).Errorf("query train station price failed, err: %v", err)
		return nil, err
	}
	priceMap := make(map[int]int)
	if err = sonic.UnmarshalString(loadData, &priceMap); err != nil {
		log.WithContext(ctx).Errorf("unmarshal train station price failed, err: %v", err)
		return nil, err
	}
	queryResp, err := configs.UserServiceClient.ListPassengerQueryByIds(ctx, &user_service.ListPassengerByIdsReq{
		Ids:      passengerIds,
		Username: req.GetUsername(),
	})
	if err != nil {
		return nil, err
	}
	passengerMap := make(map[string]*user_service.PassengerActualData, len(passengers))
	for _, passenger := range queryResp.GetData() {
		passengerMap[passenger.GetId()] = passenger
	}
	for _, item := range resp {
		if passenger, ok := passengerMap[item.PassengerId]; ok {
			item.IdType = int(passenger.GetIdType())
			item.IdCard = passenger.GetIdCard()
			item.RealName = passenger.GetRealName()
			item.Phone = passenger.GetPhone()
			item.PassengerId = passenger.GetId()
			item.Amount = priceMap[item.SeatType]
			item.DepartureDate = req.GetDepartureDate()
		}
	}
	if err = s.seatService.LockSeat(ctx, req.GetTrainId(), req.GetDeparture(), req.GetArrival(), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *TicketService) CancelTicketOrder(ctx context.Context, req *ticket_service.CancelTicketRequest) (*ticket_service.CancelTicketResponse, error) {
	resp := &ticket_service.CancelTicketResponse{}
	condition := map[string]interface{}{
		"id = ?": req.GetTicketID(),
	}
	tickets, err := dao.QueryTicket(ctx, condition)
	if err != nil {
		resp.Message = "系统错误"
		return resp, nil
	}
	if len(tickets) == 0 {
		resp.Message = "车票不存在"
		return resp, nil
	}
	ticket := tickets[0]
	var trainPurchaseTicketResults []*TrainPurchaseTicketDTO
	trainPurchaseTicketResults = append(trainPurchaseTicketResults, &TrainPurchaseTicketDTO{
		DepartureDate:  tools.ConvertTimeToStr(ticket.DepartureTime, false),
		CarriageNumber: ticket.CarriageNumber,
		SeatNumber:     ticket.SeatNumber,
	})
	// 开启事务
	tx := configs.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	tx.Begin()
	if err = s.seatService.Unlock(ctx, strconv.Itoa(int(ticket.TrainID)), ticket.Departure, ticket.Arrival, trainPurchaseTicketResults); err != nil {
		resp.Message = "系统错误"
	}
	if err = dao.CancelTicket(ctx, ticket.ID); err != nil {
		resp.Message = "系统错误"
	}
	tx.Commit()
	resp.Success = true
	return resp, nil
}

func (s *TicketService) GetTicket(ctx context.Context, req *ticket_service.GetTicketRequest) (*ticket_service.GetTicketResponse, error) {
	resp := &ticket_service.GetTicketResponse{}
	passengerResp, err := configs.UserServiceClient.ListPassengerQueryByUsername(ctx, &user_service.ListPassengerByUsernameReq{
		Username: req.GetUsername(),
	})
	if err != nil {
		resp.Message = "查询乘客信息失败"
		return resp, nil
	}
	passengers := passengerResp.GetData()
	log.WithContext(ctx).Infof("passengers: %+v", tools.MustJson(passengers))
	passengerIdMap := make(map[int64]*user_service.Passenger, len(passengers))
	passengerIds := make([]int64, 0, len(passengers))
	for _, passenger := range passengers {
		pid, _ := strconv.ParseInt(passenger.GetId(), 10, 64)
		passengerIdMap[pid] = passenger
		passengerIds = append(passengerIds, pid)
	}
	condition := map[string]interface{}{
		"passenger_id in ?":  passengerIds,
		"departure_time > ?": time.Now(),
	}
	tickets, err := dao.QueryTicket(ctx, condition)
	if err != nil {
		resp.Message = "查询车票信息失败"
		return resp, nil
	}
	data := make([]*ticket_service.MyTicketInfo, 0, len(tickets))
	for _, ticket := range tickets {
		passenger := passengerIdMap[ticket.PassengerID]
		if passenger == nil {
			continue
		}
		data = append(data, &ticket_service.MyTicketInfo{
			Id:             strconv.Itoa(int(ticket.ID)),
			TrainNumber:    ticket.TrainNumber,
			Departure:      ticket.Departure,
			Arrival:        ticket.Arrival,
			DepartureTime:  tools.ConvertTimeToStr(ticket.DepartureTime, true),
			ArrivalTime:    tools.ConvertTimeToStr(ticket.ArrivalTime, true),
			CarriageNumber: ticket.CarriageNumber,
			SeatNumber:     ticket.SeatNumber,
			PassengerName:  passenger.GetRealName(),
			TicketStatus:   int32(ticket.TicketStatus),
		})
	}
	resp.Data = data
	resp.Success = true
	return resp, nil
}
