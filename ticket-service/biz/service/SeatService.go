package service

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/tools"
	"strconv"
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
func (s *SeatService) ListSeatRemainingTicket(ctx context.Context, trainId string, departure string, arrival, departureDate string, trainCarriageList []string) ([]int, error) {
	trainID, _ := strconv.ParseInt(trainId, 10, 64)
	seatDO := &model.SeatDO{
		TrainID:       trainID,
		StartStation:  departure,
		EndStation:    arrival,
		DepartureDate: departureDate,
	}
	return dao.ListSeatRemainingTicket(ctx, seatDO, trainCarriageList)
}

// ListUsableCarriageNumber 查询列车批次有余票的车厢号集合
func (s *SeatService) ListUsableCarriageNumber(ctx context.Context, trainId int64, seatType int, departure string, arrival string, departureDate string) ([]string, error) {
	conditions := map[string]interface{}{
		"train_id = ?":       trainId,
		"seat_type = ?":      seatType,
		"start_station = ?":  departure,
		"end_station = ?":    arrival,
		"seat_status = ?":    model.SeatAvailable,
		"departure_date = ?": departureDate,
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

func (s *SeatService) updateSeatStatus(ctx context.Context, trainId int64, departure string, arrival string, trainPurchaseTicketResults []*TrainPurchaseTicketDTO, seatStatus int32) error {
	// 改成事务
	routes, err := s.trainStationService.ListTakeoutTrainStationRoute(ctx, trainId, departure, arrival)
	if err != nil {
		log.WithContext(ctx).Errorf("list takeout train station route failed, err: %v", err)
		return err
	}
	log.WithContext(ctx).Infof("takeout routes: %v", tools.MustJson(routes))
	seats := make([]*model.SeatDO, 0)
	for _, item := range trainPurchaseTicketResults {
		for _, route := range routes {
			seats = append(seats, &model.SeatDO{
				TrainID:        trainId,
				CarriageNumber: item.CarriageNumber,
				SeatNumber:     item.SeatNumber,
				StartStation:   route.StartStation,
				EndStation:     route.EndStation,
				DepartureDate:  item.DepartureDate,
				SeatStatus:     int(seatStatus),
			})
		}
	}
	eqCondition := map[string]interface{}{
		"seat_status": 1 - seatStatus,
	}
	log.WithContext(ctx).Infof("update seat status: %v", tools.MustJson(seats))
	return dao.UpsertSeats(ctx, seats, false, []string{"seat_status"}, eqCondition)
}

// LockSeat 锁定选中以及沿途车票状态
func (s *SeatService) LockSeat(ctx context.Context, trainId string, departure string, arrival string, trainPurchaseTicketRespList []*TrainPurchaseTicketDTO) error {
	tid, _ := strconv.ParseInt(trainId, 10, 64)
	return s.updateSeatStatus(ctx, tid, departure, arrival, trainPurchaseTicketRespList, model.SeatLocked)
}

// Unlock 解锁选中以及沿途车票状态
func (s *SeatService) Unlock(ctx context.Context, trainId string, departure string, arrival string, trainPurchaseTicketResults []*TrainPurchaseTicketDTO) error {
	tid, _ := strconv.ParseInt(trainId, 10, 64)
	return s.updateSeatStatus(ctx, tid, departure, arrival, trainPurchaseTicketResults, model.SeatAvailable)
}

type TrainPurchaseTicketDTO struct {
	DepartureDate  string `json:"departureDate"`  // 出发日期
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
