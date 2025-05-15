package dao

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
	"log"
)

func QuerySeats(ctx context.Context, condition map[string]interface{}) ([]*model.SeatDO, error) {
	var seats []*model.SeatDO
	query := configs.DB.Model(&model.SeatDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	err := query.Scan(&seats).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query seat failed, err: %v", err)
		return nil, err
	}
	return seats, nil
}

func UpsertSeats(ctx context.Context, seat []*model.SeatDO, ignoreDup bool, updateCols []string, eqCondition map[string]interface{}) (int64, error) {
	// 使用clause
	var clauses []clause.Expression
	onConflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "train_id"}, {Name: "carriage_number"}, {Name: "seat_number"},
			{Name: "start_station"}, {Name: "end_station"}, {Name: "departure_date"}},
	}
	onConflict.DoUpdates = clause.AssignmentColumns(updateCols)

	if ignoreDup {
		onConflict.DoNothing = true
	} else {
		onConflict.DoUpdates = clause.AssignmentColumns(updateCols)
	}
	clauses = append(clauses, onConflict)
	if len(eqCondition) > 0 {
		where := clause.Where{}
		for k, v := range eqCondition {
			where.Exprs = append(where.Exprs, clause.Eq{Column: clause.Column{Name: k}, Value: v})
		}
		clauses = append(clauses, where)
	}
	// 执行 upsert 并获取结果
	tx := configs.DB.Model(&model.SeatDO{}).WithContext(ctx).Clauses(clauses...).Create(&seat)
	return tx.RowsAffected, tx.Error
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
		Where("departure_date = ?", seatDO.DepartureDate).
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

type SeatCountByType struct {
	SeatType int   `json:"seat_type"`
	Count    int64 `json:"count"`
}

func QuerySeatCountGroupedByType(ctx context.Context, trainID int64, startStation, endStation, departureDate string) ([]SeatCountByType, error) {
	var result []SeatCountByType

	err := configs.DB.
		Model(&model.SeatDO{}).
		Select("seat_type, COUNT(*) as count").
		Where("train_id = ? AND start_station = ? AND end_station = ? AND departure_date = ? AND seat_status = 0",
			trainID, startStation, endStation, departureDate).
		Group("seat_type").
		Scan(&result).Error

	return result, err
}

// BatchInsertSeat 批量插入座位
func BatchInsertSeat(ctx context.Context, seats []*model.SeatDO) error {
	err := configs.DB.WithContext(ctx).Create(&seats).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			log.Println("唯一索引冲突，跳过插入")
		} else {
			return err
		}
	}
	return nil
}

// InsertSeat 批量插入座位
func InsertSeat(ctx context.Context, seats *model.SeatDO) error {
	return configs.DB.WithContext(ctx).Save(&seats).Error
}
