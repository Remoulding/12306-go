package script

import (
	"context"
	"errors"
	"fmt"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func init() {
	// 初始化数据库
	configs.InitDBInstance()

	seatMapper = make(map[int]*SeatInfo)
	seatMapper[0] = &SeatInfo{
		rowNum:      2,
		colRune:     []rune{'A', 'C', 'F'},
		CarriageNum: []string{"01", "16"},
	}
	seatMapper[1] = &SeatInfo{
		rowNum:      7,
		colRune:     []rune{'A', 'C', 'D', 'F'},
		CarriageNum: []string{"02", "03", "04", "05", "06", "07"},
	}
	seatMapper[2] = &SeatInfo{
		rowNum:      18,
		colRune:     []rune{'A', 'B', 'C', 'D', 'F'},
		CarriageNum: []string{"08", "09", "10", "11", "12", "13", "14", "15"},
	}
}

type SeatInfo struct {
	rowNum      int
	CarriageNum []string
	colRune     []rune
}

var departureDates []string

var seatMapper map[int]*SeatInfo

func InitData() {
	departureDates = make([]string, 14)
	now := time.Now()
	for i := range departureDates {
		departureDates[i] = now.AddDate(0, 0, i).Format("2006-01-02")
	}
	// 读取所有的train
	ctx := context.Background()
	trains, _ := dao.QueryTrain(ctx, nil)

	for _, train := range trains {
		// 批量插入座位
		_ = batchInsertSeats(ctx, train.ID)

	}
}

func batchInsertSeats(ctx context.Context, trainID int64) error {
	// 批量插入座位
	exp := map[string]interface{}{}
	exp["train_id = ?"] = trainID
	relation, err := dao.QueryTrainStationRelation(ctx, exp)
	if err != nil {
		return err
	}
	relationPrice, err := dao.QueryTrainStationPrice(ctx, map[string]interface{}{"train_id": trainID})
	if err != nil {
		return err
	}
	priceMapper := make(map[string]map[int]int)
	for _, r := range relationPrice {
		key := fmt.Sprintf("%d_%s_%s", r.TrainID, r.Departure, r.Arrival)
		if _, ok := priceMapper[key]; !ok {
			priceMapper[key] = make(map[int]int)
		}
		priceMapper[key][r.SeatType] = r.Price
	}
	for _, r := range relation {
		key := fmt.Sprintf("%d_%s_%s", r.TrainID, r.Departure, r.Arrival)
		if _, ok := priceMapper[key]; !ok {
			continue
		}
		for _, date := range departureDates {
			for seatType, price := range priceMapper[key] {
				err = insertSeat(ctx, trainID, seatType, r.Departure, r.Arrival, date, price)
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					continue
				}
				if err != nil {
					fmt.Printf("insert seat failed, err: %v\n", err)
					return err
				}
			}
		}
	}

	return nil
}

func insertSeat(ctx context.Context, trainID int64, seatType int, startStation, endStation string, departureDate string, price int) error {
	// 插入座位
	for _, carriageNum := range seatMapper[seatType].CarriageNum {
		seatDOS := make([]*model.SeatDO, 0)
		for i := 1; i <= seatMapper[seatType].rowNum; i++ {
			for _, col := range seatMapper[seatType].colRune {
				seatStr := strconv.Itoa(i) + string(col)
				if i < 10 {
					seatStr = "0" + seatStr
				}
				seatDO := &model.SeatDO{
					TrainID:        trainID,
					CarriageNumber: carriageNum,
					SeatNumber:     seatStr,
					SeatType:       seatType,
					StartStation:   startStation,
					EndStation:     endStation,
					DepartureDate:  departureDate,
					SeatStatus:     model.SeatAvailable,
					Price:          price,
				}
				seatDOS = append(seatDOS, seatDO)
			}
		}
		//err := dao.BatchInsertSeat(ctx, seatDOS)
		_, err := dao.UpsertSeats(ctx, seatDOS, true, nil, nil)
		if err != nil {
			return err
		}
	}
	return nil
}
