package service

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"strconv"
	"sync"
)

type SeatService struct {
	trainStationService *TrainStationService
}

func NewSeatService() *SeatService {
	return &SeatService{
		trainStationService: NewTrainStationService(),
	}
}

// ListAvailableSeat 获取列车车厢中可用的座位集合
func (s *SeatService) ListAvailableSeat(ctx context.Context, trainId string, carriageNumber string, seatType int, departure string, arrival string) ([]string, error) {
	conditions := map[string]interface{}{
		"train_id = ?":        trainId,
		"carriage_number = ?": carriageNumber,
		"seat_type = ?":       seatType,
		"start_station = ?":   departure,
		"end_station = ?":     arrival,
		"seat_status = ?":     model.SeatAvailable,
	}
	seats, err := dao.QuerySeats(ctx, conditions)
	if err != nil {
		log.WithContext(ctx).Errorf("query seat failed, err: %v", err)
		return nil, err
	}
	var seatNumbers []string
	for _, seat := range seats {
		seatNumbers = append(seatNumbers, seat.SeatNumber)
	}
	return seatNumbers, nil
}

// ListSeatRemainingTicket 获取列车车厢余票集合
func (s *SeatService) ListSeatRemainingTicket(ctx context.Context, trainId string, departure string, arrival string, trainCarriageList []string) ([]int, error) {
	trainID, _ := strconv.ParseInt(trainId, 10, 64)
	seatDO := &model.SeatDO{
		TrainID:      trainID,
		StartStation: departure,
		EndStation:   arrival,
	}
	return dao.ListSeatRemainingTicket(ctx, seatDO, trainCarriageList)
}

// ListUsableCarriageNumber 查询列车批次有余票的车厢号集合
func (s *SeatService) ListUsableCarriageNumber(ctx context.Context, trainId string, seatType int, departure string, arrival string) ([]string, error) {
	conditions := map[string]interface{}{
		"train_id = ?":      trainId,
		"seat_type = ?":     seatType,
		"start_station = ?": departure,
		"end_station = ?":   arrival,
		"seat_status = ?":   model.SeatAvailable,
	}
	seats, err := dao.QuerySeats(ctx, conditions)
	if err != nil {
		log.WithContext(ctx).Errorf("query seat failed, err: %v", err)
		return nil, err
	}
	var carriageNumbers []string
	var carriageNumberMap = make(map[string]struct{})
	for _, seat := range seats {
		carriageNumberMap[seat.CarriageNumber] = struct{}{}
	}
	for carriageNumber := range carriageNumberMap {
		carriageNumbers = append(carriageNumbers, carriageNumber)
	}
	return carriageNumbers, nil
}

func (s *SeatService) updateSeatStatus(ctx context.Context, trainId string, departure string, arrival string, trainPurchaseTicketResults []*TrainPurchaseTicketDTO, seatStatus int32) error {
	routes, err := s.trainStationService.ListTakeoutTrainStationRoute(ctx, trainId, departure, arrival)
	if err != nil {
		log.WithContext(ctx).Errorf("list takeout train station route failed, err: %v", err)
		return err
	}
	var wg sync.WaitGroup
	errs := make(chan error, len(trainPurchaseTicketResults)*len(routes))
	for _, item := range trainPurchaseTicketResults {
		for _, route := range routes {
			wg.Add(1)
			go func(item *TrainPurchaseTicketDTO, route *Route) {
				defer wg.Done()
				condition := map[string]interface{}{
					"train_id = ?":        trainId,
					"carriage_number = ?": item.CarriageNumber,
					"seat_number = ?":     item.SeatNumber,
					"start_station = ?":   route.StartStation,
					"end_station = ?":     route.EndStation,
				}
				update := map[string]interface{}{
					"seat_status": seatStatus,
				}
				errs <- dao.UpdateSeats(ctx, condition, update)
			}(item, route) // 显式传递 item 和 route，防止变量捕获问题
		}
	}

	// 另外起一个 goroutine 关闭 errs，防止主 goroutine 死锁
	go func() {
		wg.Wait()
		close(errs)
	}()

	// 收集错误
	for err := range errs {
		if err != nil {
			log.WithContext(ctx).Errorf("update seat failed, err: %v", err)
			return err
		}
	}
	return nil
}

// LockSeat 锁定选中以及沿途车票状态
func (s *SeatService) LockSeat(ctx context.Context, trainId string, departure string, arrival string, trainPurchaseTicketRespList []*TrainPurchaseTicketDTO) error {
	return s.updateSeatStatus(ctx, trainId, departure, arrival, trainPurchaseTicketRespList, model.SeatLocked)
}

// Unlock 解锁选中以及沿途车票状态
func (s *SeatService) Unlock(ctx context.Context, trainId string, departure string, arrival string, trainPurchaseTicketResults []*TrainPurchaseTicketDTO) error {
	return s.updateSeatStatus(ctx, trainId, departure, arrival, trainPurchaseTicketResults, model.SeatAvailable)
}

type TrainPurchaseTicketDTO struct {
	PassengerId    string `json:"passengerId"`    // 乘车人 ID
	RealName       string `json:"realName"`       // 乘车人姓名
	IdType         int    `json:"idType"`         // 乘车人证件类型
	IdCard         string `json:"idCard"`         // 乘车人证件号
	Phone          string `json:"phone"`          // 乘车人手机号
	UserType       int    `json:"userType"`       // 用户类型 0：成人 1：儿童 2：学生 3：残疾军人
	SeatType       int    `json:"seatType"`       // 席别类型
	CarriageNumber string `json:"carriageNumber"` // 车厢号
	SeatNumber     string `json:"seatNumber"`     // 座位号
	Amount         int    `json:"amount"`         // 座位金额
}
