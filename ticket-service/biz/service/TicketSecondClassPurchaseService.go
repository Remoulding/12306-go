package service

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/samber/lo"
	"strconv"
)

type TicketSecondClassPurchaseService struct {
	ctx         context.Context
	seatType    int
	SeatService *SeatService
}

func NewTicketSecondClassPurchaseService(ctx context.Context) *TicketSecondClassPurchaseService {
	return &TicketSecondClassPurchaseService{
		ctx:         ctx,
		seatType:    2,
		SeatService: NewSeatService(),
	}
}

// Pair 结构体
type Pair struct {
	X, Y int
}

// SEAT_Y_INT 模拟座位列字符到整数的映射
var SEAT_Y_INT = map[rune]int{
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4,
}

func (s *TicketSecondClassPurchaseService) SelectSeats(passengerList []*ticket_service.PurchaseTicketPassengerInfo,
	req *ticket_service.PurchaseTicketRequest) ([]*TrainPurchaseTicketDTO, error) {
	trainId := req.GetTrainId()
	departure := req.GetDeparture()
	arrival := req.GetArrival()
	carriageNumber, err := s.SeatService.ListUsableCarriageNumber(s.ctx, trainId, s.seatType, departure, arrival)
	if err != nil {
		return nil, err
	}
	remainingTicket, err := s.SeatService.ListSeatRemainingTicket(s.ctx, trainId, departure, arrival, carriageNumber)
	if err != nil {
		return nil, err
	}
	if lo.Sum(remainingTicket) < len(passengerList) {
		return nil, errors.New("余票不足")
	}
	if len(req.GetChooseSeats()) > 0 {
		return s.findMatchSeats(req, carriageNumber, remainingTicket)
	}
	return s.selectSeats(req, carriageNumber, remainingTicket)
}

func (s *TicketSecondClassPurchaseService) selectSeats(req *ticket_service.PurchaseTicketRequest, carriageNumber []string,
	remainingTickets []int) ([]*TrainPurchaseTicketDTO, error) {
	resp := make([]*TrainPurchaseTicketDTO, 0, len(req.GetPassengers()))

	return resp, nil
}

func (s *TicketSecondClassPurchaseService) findMatchSeats(req *ticket_service.PurchaseTicketRequest, carriageNumber []string,
	remainingTickets []int) ([]*TrainPurchaseTicketDTO, error) {
	resp := make([]*TrainPurchaseTicketDTO, 0, len(req.GetPassengers()))
	carriageSeatMap := make(map[string][]*Pair)
	passengerNum := len(req.GetPassengers())
	for i := range remainingTickets {
		carriageNumb := carriageNumber[i]
		listAvailableSeat, err := s.SeatService.ListAvailableSeat(s.ctx, req.GetTrainId(), carriageNumb, s.seatType, req.GetDeparture(), req.GetArrival())
		if err != nil {
			return nil, err
		}
		listAvailableSeatSet := make(map[string]struct{})
		for _, seat := range listAvailableSeat {
			listAvailableSeatSet[seat] = struct{}{}
		}
		actualSeats := make([][]int, 18)
		vacantSeats := make([]*Pair, 0) // 未被占用的座位

		// 处理实际座位数据
		for j := 1; j <= 18; j++ {
			actualSeats[j-1] = make([]int, 5)
			for k := 1; k <= 5; k++ {
				seatStr := strconv.Itoa(j) + Convert(s.seatType, k)
				if j < 10 {
					seatStr = "0" + seatStr
				}
				if _, ok := listAvailableSeatSet[seatStr]; ok {
					actualSeats[j-1][k-1] = 0
				} else {
					actualSeats[j-1][k-1] = 1
				}
				if actualSeats[j-1][k-1] == 0 {
					vacantSeats = append(vacantSeats, &Pair{j - 1, k - 1})
				}
			}
		}
		// 用户选择的座位
		sureSeatList := s.calcChooseSeatLevelPairList(actualSeats, req.GetChooseSeats())
		selectSeats := make([]string, 0, len(req.GetPassengers()))
		if len(sureSeatList) > 0 && len(vacantSeats) >= len(req.GetPassengers()) {
			vacantSeatList := make([]*Pair, 0)
			if len(sureSeatList) != len(req.GetPassengers()) {
				// 补充座位
				for _, sureSeat := range sureSeatList {
					actualSeats[sureSeat.X][sureSeat.Y] = 1
				}
			}
			for i1 := 0; i1 < 18; i1++ {
				for j := 0; j < 5; j++ {
					if actualSeats[i1][j] == 0 {
						vacantSeatList = append(vacantSeatList, &Pair{i1, j})
					}
				}
			}
			needAdd := len(req.GetPassengers()) - len(sureSeatList)
			sureSeatList = append(sureSeatList, vacantSeatList[:needAdd]...)
			s.addSeats(selectSeats, sureSeatList)

			for i1, selectSeat := range selectSeats {
				passengerInfo := req.GetPassengers()[i1]
				resp = append(resp, &TrainPurchaseTicketDTO{
					PassengerId:    passengerInfo.GetPassengerId(),
					CarriageNumber: carriageNumb,
					SeatNumber:     selectSeat,
					SeatType:       int(passengerInfo.GetSeatType()),
				})
			}
			return resp, nil
		}
		// 降级分配
		if len(vacantSeats) > 0 {
			carriageSeatMap[carriageNumb] = vacantSeats
			if i != len(remainingTickets)-1 {
				continue
			}
			var findCarriageSureSeats []*Pair
			var findCarriageNum string
			for carriageNum, carriageSeats := range carriageSeatMap {
				if len(carriageSeats) >= len(req.GetPassengers()) {
					findCarriageSureSeats = carriageSeats
					findCarriageNum = carriageNum
					break
				}
			}
			// 单独一个车厢
			if len(findCarriageSureSeats) > 0 {
				needSeats := findCarriageSureSeats[:len(req.GetPassengers())]
				s.addSeats(selectSeats, needSeats)
				for i1, selectSeat := range selectSeats {
					passengerInfo := req.GetPassengers()[i1]
					resp = append(resp, &TrainPurchaseTicketDTO{
						PassengerId:    passengerInfo.GetPassengerId(),
						CarriageNumber: findCarriageNum,
						SeatNumber:     selectSeat,
						SeatType:       int(passengerInfo.GetSeatType()),
					})
				}
				return resp, nil
			} else {
				// 多个车厢
				leftSeatCnt := 0
				for _, carriageSeats := range carriageSeatMap {
					leftSeatCnt += len(carriageSeats)
				}
				if leftSeatCnt < len(req.GetPassengers()) {
					return nil, errors.New("余票不足")
				}
				currentSeatSize := 0
				for _, carriageSeats := range carriageSeatMap {
					if len(carriageSeats)+currentSeatSize < passengerNum {
						currentSeatSize += len(carriageSeats)
						s.addSeats(selectSeats, carriageSeats)
					} else {
						needSeats := carriageSeats[:passengerNum-currentSeatSize]
						s.addSeats(selectSeats, needSeats)
						for i1, selectSeat := range selectSeats {
							passengerInfo := req.GetPassengers()[i1]
							resp = append(resp, &TrainPurchaseTicketDTO{
								PassengerId:    passengerInfo.GetPassengerId(),
								CarriageNumber: findCarriageNum,
								SeatNumber:     selectSeat,
								SeatType:       int(passengerInfo.GetSeatType()),
							})
						}
						return resp, nil
					}
				}
			}

		}
	}
	return resp, nil
}

func (s *TicketSecondClassPurchaseService) addSeats(targetSeats []string, seatPairs []*Pair) {
	for _, sureSeat := range seatPairs {
		seatStr := strconv.Itoa(sureSeat.X+1) + Convert(s.seatType, sureSeat.Y+1)
		if sureSeat.X < 9 {
			seatStr = "0" + seatStr
		}
		targetSeats = append(targetSeats, seatStr)
	}
}

// calcChooseSeatLevelPairList 计算可用的座位列表
func (s *TicketSecondClassPurchaseService) calcChooseSeatLevelPairList(actualSeats [][]int, chooseSeatList []string) []*Pair {
	if len(chooseSeatList) == 0 {
		return nil
	}

	// 解析第一个选座
	firstChooseSeat := chooseSeatList[0]
	firstSeatX, _ := strconv.Atoi(firstChooseSeat[1:])
	firstSeatY := SEAT_Y_INT[rune(firstChooseSeat[0])]

	chooseSeatLevelPairList := []*Pair{{firstSeatX, firstSeatY}}
	minLevelX := 0

	// 计算座位相对偏移
	for i := 1; i < len(chooseSeatList); i++ {
		chooseSeat := chooseSeatList[i]
		chooseSeatX, _ := strconv.Atoi(chooseSeat[1:])
		chooseSeatY := SEAT_Y_INT[rune(chooseSeat[0])]

		minLevelX = min(minLevelX, chooseSeatX-firstSeatX)
		chooseSeatLevelPairList = append(chooseSeatLevelPairList, &Pair{chooseSeatX - firstSeatX, chooseSeatY - firstSeatY})
	}

	// 遍历 actualSeats 查找可用座位
	for i := abs(minLevelX); i < len(actualSeats); i++ {
		var sureSeatList []*Pair
		if actualSeats[i][firstSeatY] == 0 {
			sureSeatList = append(sureSeatList, &Pair{i, firstSeatY})

			for j := 1; j < len(chooseSeatList); j++ {
				pair := chooseSeatLevelPairList[j]
				x := i + pair.X
				y := firstSeatY + pair.Y

				// 超出边界
				if x >= len(actualSeats) {
					return nil
				}

				if actualSeats[x][y] == 0 {
					sureSeatList = append(sureSeatList, &Pair{x, y})
				} else {
					break
				}
			}
		}

		if len(sureSeatList) == len(chooseSeatList) {
			return sureSeatList
		}
	}

	return nil
}

var (
	trainBusinessClassSeatNumberMap = map[int]string{
		1: "A",
		2: "C",
		3: "F",
	}

	trainFirstClassSeatNumberMap = map[int]string{
		1: "A",
		2: "C",
		3: "D",
		4: "F",
	}

	trainSecondClassSeatNumberMap = map[int]string{
		1: "A",
		2: "B",
		3: "C",
		4: "D",
		5: "F",
	}
)

// Convert 根据座位类型转换座位号
func Convert(seatType, num int) string {
	switch seatType {
	case 0:
		return trainBusinessClassSeatNumberMap[num]
	case 1:
		return trainFirstClassSeatNumberMap[num]
	case 2:
		return trainSecondClassSeatNumberMap[num]
	default:
		return ""
	}
}

// 辅助函数：获取最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 辅助函数：取绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
