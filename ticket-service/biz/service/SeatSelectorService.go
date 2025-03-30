package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/samber/lo"
	"strconv"
)

type SeatSelectorService struct {
	ctx         context.Context
	seatType    int
	seatService *SeatService
	seatYInt    map[rune]int
	rowSize     int
	colSize     int
}

func NewSeatSelectorService(ctx context.Context, seatType int) *SeatSelectorService {
	s := &SeatSelectorService{
		ctx:         ctx,
		seatType:    seatType,
		seatService: NewSeatService(),
	}
	if seatType == 0 {
		s.seatYInt = map[rune]int{
			'A': 0, 'C': 1, 'F': 2,
		}
		s.rowSize = 2
		s.colSize = 3
	}
	if seatType == 1 {
		s.seatYInt = map[rune]int{
			'A': 0, 'C': 1, 'D': 2, 'F': 3,
		}
		s.rowSize = 7
		s.colSize = 4
	}
	if seatType == 2 {
		s.seatYInt = map[rune]int{
			'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4,
		}
		s.rowSize = 18
		s.colSize = 5
	}
	return s
}

// Pair 结构体
type Pair struct {
	X, Y int
}

func (s *SeatSelectorService) SelectSeats(passengerList []*ticket_service.PurchaseTicketPassengerInfo,
	req *ticket_service.PurchaseTicketRequest) ([]*TrainPurchaseTicketDTO, error) {
	trainId := req.GetTrainId()
	departure := req.GetDeparture()
	arrival := req.GetArrival()
	carriageNumber, err := s.seatService.ListUsableCarriageNumber(s.ctx, trainId, s.seatType, departure, arrival, req.GetDepartureDate())
	if err != nil {
		return nil, err
	}
	remainingTicket, err := s.seatService.ListSeatRemainingTicket(s.ctx, trainId, departure, arrival, req.GetDepartureDate(), carriageNumber)
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

func (s *SeatSelectorService) selectSeats(req *ticket_service.PurchaseTicketRequest, carriageNumber []string,
	remainingTickets []int) ([]*TrainPurchaseTicketDTO, error) {
	resp := make([]*TrainPurchaseTicketDTO, 0, len(req.GetPassengers()))
	passengerSize := len(req.GetPassengers())
	demotionStockNumMap := make(map[string]int)
	actualSeatsMap := make(map[string][][]int)
	carriagesNumberSeatsMap := make(map[string][][]int)
	for i := range remainingTickets {
		carriageNumb := carriageNumber[i]
		listAvailableSeat, err := s.seatService.ListAvailableSeat(s.ctx, req.GetTrainId(), carriageNumb, s.seatType, req.GetDeparture(), req.GetArrival())
		if err != nil {
			return nil, err
		}
		listAvailableSeatSet := make(map[string]struct{})
		for _, seat := range listAvailableSeat {
			listAvailableSeatSet[seat] = struct{}{}
		}

		actualSeats := make([][]int, s.rowSize)
		// 处理实际座位数据
		for j := 1; j <= s.rowSize; j++ {
			actualSeats[j-1] = make([]int, s.colSize)
			for k := 1; k <= s.colSize; k++ {
				seatStr := strconv.Itoa(j) + Convert(s.seatType, k)
				if j < 10 {
					seatStr = "0" + seatStr
				}
				if _, ok := listAvailableSeatSet[seatStr]; ok {
					actualSeats[j-1][k-1] = 0
				} else {
					actualSeats[j-1][k-1] = 1
				}
			}
		}

		selectSeats := SelectAdjacentSeats(passengerSize, actualSeats)
		// 有连坐
		if len(selectSeats) != 0 {
			carriagesNumberSeatsMap[carriageNumb] = selectSeats
			break
		}
		// 无连坐，计算余票
		demotionStockNum := 0
		for _, seats := range actualSeats {
			for _, seat := range seats {
				if seat == 0 {
					demotionStockNum++
				}
			}
		}
		actualSeatsMap[carriageNumb] = actualSeats
		demotionStockNumMap[carriageNumb] = demotionStockNum
		if i < len(remainingTickets)-1 {
			continue
		}
		// 同车厢
		for carrNum, carrSeatNum := range demotionStockNumMap {
			if carrSeatNum >= passengerSize {
				selectSeats = SelectNonAdjacentSeats(passengerSize, actualSeatsMap[carrNum])
				if len(selectSeats) == passengerSize {
					carriagesNumberSeatsMap[carrNum] = selectSeats
					break
				}
			}
		}

		// 不同车厢
		selectSeatCnt := 0
		for carrNum, carrSeatNum := range demotionStockNumMap {
			if selectSeatCnt < passengerSize {
				addSeatCnt := min(carrSeatNum, passengerSize-selectSeatCnt)
				selectSeatCnt += addSeatCnt
				carriagesNumberSeatsMap[carrNum] = SelectNonAdjacentSeats(addSeatCnt, actualSeatsMap[carrNum])
			} else {
				break
			}
		}

		if selectSeatCnt < passengerSize {
			return nil, errors.New("余票不足")
		}
	}
	for carrNum, carrSeats := range carriagesNumberSeatsMap {
		selectSeats := make([]string, 0, passengerSize)
		for _, seat := range carrSeats {
			seatStr := strconv.Itoa(seat[0]) + Convert(s.seatType, seat[1])
			if seat[0] < 10 {
				seatStr = "0" + seatStr
			}
			selectSeats = append(selectSeats, seatStr)
		}
		for i, selectSeat := range selectSeats {
			passengerInfo := req.GetPassengers()[i]
			resp = append(resp, &TrainPurchaseTicketDTO{
				PassengerId:    passengerInfo.GetPassengerId(),
				CarriageNumber: carrNum,
				SeatNumber:     selectSeat,
				SeatType:       int(passengerInfo.GetSeatType()),
			})
		}
	}
	return resp, nil
}

func (s *SeatSelectorService) findMatchSeats(req *ticket_service.PurchaseTicketRequest, carriageNumber []string,
	remainingTickets []int) ([]*TrainPurchaseTicketDTO, error) {
	resp := make([]*TrainPurchaseTicketDTO, 0, len(req.GetPassengers()))
	carriageSeatMap := make(map[string][]*Pair)
	passengerNum := len(req.GetPassengers())
	for i := range remainingTickets {
		carriageNumb := carriageNumber[i]
		listAvailableSeat, err := s.seatService.ListAvailableSeat(s.ctx, req.GetTrainId(), carriageNumb, s.seatType, req.GetDeparture(), req.GetArrival())
		if err != nil {
			return nil, err
		}
		listAvailableSeatSet := make(map[string]struct{})
		for _, seat := range listAvailableSeat {
			listAvailableSeatSet[seat] = struct{}{}
		}
		actualSeats := make([][]int, s.rowSize)
		vacantSeats := make([]*Pair, 0) // 未被占用的座位

		// 处理实际座位数据
		for j := 1; j <= s.rowSize; j++ {
			actualSeats[j-1] = make([]int, s.colSize)
			for k := 1; k <= s.colSize; k++ {
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
			for i1 := 0; i1 < s.rowSize; i1++ {
				for j := 0; j < s.colSize; j++ {
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

func (s *SeatSelectorService) addSeats(targetSeats []string, seatPairs []*Pair) {
	for _, sureSeat := range seatPairs {
		seatStr := strconv.Itoa(sureSeat.X+1) + Convert(s.seatType, sureSeat.Y+1)
		if sureSeat.X < 9 {
			seatStr = "0" + seatStr
		}
		targetSeats = append(targetSeats, seatStr)
	}
}

// calcChooseSeatLevelPairList 计算可用的座位列表
func (s *SeatSelectorService) calcChooseSeatLevelPairList(actualSeats [][]int, chooseSeatList []string) []*Pair {
	if len(chooseSeatList) == 0 {
		return nil
	}

	// 解析第一个选座
	firstChooseSeat := chooseSeatList[0]
	firstSeatX, _ := strconv.Atoi(firstChooseSeat[1:])
	firstSeatY := s.seatYInt[rune(firstChooseSeat[0])]

	chooseSeatLevelPairList := []*Pair{{firstSeatX, firstSeatY}}
	minLevelX := 0

	// 计算座位相对偏移
	for i := 1; i < len(chooseSeatList); i++ {
		chooseSeat := chooseSeatList[i]
		chooseSeatX, _ := strconv.Atoi(chooseSeat[1:])
		chooseSeatY := s.seatYInt[rune(chooseSeat[0])]

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

// SelectAdjacentSeats 查找相邻座位
// 参数：numSeats 需要选择的座位数，seatLayout 座位布局矩阵（0表示可用）
// 返回值：选择的座位位置数组（行列从1开始计数），找不到时返回nil
func SelectAdjacentSeats(numSeats int, seatLayout [][]int) [][]int {
	rows := len(seatLayout)
	if rows == 0 {
		return nil
	}
	cols := len(seatLayout[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if seatLayout[i][j] == 0 {
				// 找到连续空位
				consecutive := 0
				startCol := j
				for k := j; k < cols; k++ {
					if seatLayout[i][k] == 0 {
						consecutive++
						if consecutive == numSeats {
							return convertToActualSeat(i, startCol, numSeats, true)
						}
					} else {
						consecutive = 0
						startCol = k + 1
					}
				}
			}
		}
	}
	return nil
}

// SelectNonAdjacentSeats 查找非相邻座位
func SelectNonAdjacentSeats(numSeats int, seatLayout [][]int) [][]int {
	var selected [][]int
	rows := len(seatLayout)
	if rows == 0 {
		return nil
	}
	cols := len(seatLayout[0])

outer:
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if seatLayout[i][j] == 0 {
				selected = append(selected, []int{i + 1, j + 1})
				if len(selected) == numSeats {
					break outer
				}
			}
		}
	}

	if len(selected) < numSeats {
		return nil
	}
	return selected
}

// convertToActualSeat 转换座位坐标（从0-based到1-based）
func convertToActualSeat(row, startCol, numSeats int, horizontal bool) [][]int {
	result := make([][]int, numSeats)
	for i := 0; i < numSeats; i++ {
		if horizontal {
			result[i] = []int{row + 1, startCol + i + 1}
		} else {
			// 纵向排列（当前实现仅处理横向）
			result[i] = []int{row + i + 1, startCol + 1}
		}
	}
	return result
}

// 示例用法
func ExampleUsage() {
	// 测试相邻座位选择
	seatLayout := [][]int{
		{1, 1, 1, 1},
		{1, 1, 1, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}

	if selected := SelectAdjacentSeats(2, seatLayout); selected != nil {
		fmt.Println("找到相邻座位：")
		for _, s := range selected {
			fmt.Printf("第%d排 第%d列\n", s[0], s[1])

		}
	} else {
		fmt.Println("没有足够相邻座位")
	}

	// 测试非相邻座位选择
	seatLayout2 := [][]int{
		{1, 0, 1, 1},
		{1, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 0, 0, 0},
	}

	if selected := SelectNonAdjacentSeats(3, seatLayout2); selected != nil {
		fmt.Println("\n找到非相邻座位：")
		for _, s := range selected {
			fmt.Printf("第%d排 第%d列\n", s[0], s[1])
		}
	} else {
		fmt.Println("没有足够空位")
	}
}

/*
示例输出：
找到相邻座位：
第4排 第1列
第4排 第2列

找到非相邻座位：
第1排 第2列
第2排 第3列
第2排 第4列
*/
