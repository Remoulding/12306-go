package dao

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

func QuerySeats(ctx context.Context, condition map[string]interface{}) ([]*model.SeatDO, error) {
	var seats []*model.SeatDO
	query := configs.DB.Model(&model.SeatDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	err := query.Scan(seats).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query seat failed, err: %v", err)
		return nil, err
	}
	return seats, nil
}

func UpdateSeats(ctx context.Context, condition, update map[string]interface{}) error {
	tx := configs.DB.Model(&model.SeatDO{})
	for exp, val := range condition {
		tx = tx.Where(exp, val)
	}
	for exp, val := range update {
		tx = tx.Update(exp, val)
	}
	if err := tx.Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("query seat failed, err: %v", err)
		return err
	}
	return nil
}

type SeatRemainingResult struct {
	CarriageNumber string `gorm:"column:carriage_number"`
	Count          int    `gorm:"column:count"`
}

// ListSeatRemainingTicket 查询余票数量
func ListSeatRemainingTicket(ctx context.Context, seatDO *model.SeatDO, carriageList []string) ([]int, error) {
	if seatDO == nil || len(carriageList) == 0 {
		return nil, errors.New("invalid params")
	}
	var results []SeatRemainingResult
	// 构建查询
	err := configs.DB.Model(&model.SeatDO{}).
		Select("carriage_number, COUNT(*) as count").
		Where("train_id = ?", seatDO.TrainID).
		Where("start_station = ?", seatDO.StartStation).
		Where("end_station = ?", seatDO.EndStation).
		Where("seat_status = 0").
		Where("departure_date = ", seatDO.DepartureDate).
		Where("carriage_number IN ?", carriageList).
		Group("carriage_number").
		Find(&results).Error

	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query seat remaining failed, err: %v", err)
		return nil, err
	}

	// 转换为按输入顺序的余票数列表
	resultMap := make(map[string]int)
	for _, r := range results {
		resultMap[r.CarriageNumber] = r.Count
	}

	// 保证输出顺序与输入carriageList一致
	finalCounts := make([]int, len(carriageList))
	for i, carriage := range carriageList {
		if count, exists := resultMap[carriage]; exists {
			finalCounts[i] = count
		} else {
			finalCounts[i] = 0
		}
	}
	return finalCounts, nil
}
