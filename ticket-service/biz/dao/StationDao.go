package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"gorm.io/gorm/clause"
)

func QueryStationByName(ctx context.Context, name string) (*model.StationDO, error) {
	var station *model.StationDO
	name += "%"
	err := configs.DB.Model(&model.StationDO{}).Where("name like ?", name).
		Or("spell like ?", name).First(station).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query station failed, err: %v", err)
		return nil, err
	}
	return station, nil
}

func QueryStations(ctx context.Context, condition map[string]interface{}) ([]*model.StationDO, error) {
	var stations []*model.StationDO
	query := configs.DB.Model(&model.StationDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	if err := query.Scan(&stations).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("query station failed, err: %v", err)
		return nil, err
	}
	return stations, nil
}

func UpsertStations(ctx context.Context, stations []*model.StationDO, ignoreDup bool, updateCols []string) (int64, error) {
	// 使用clause
	var clauses []clause.Expression
	onConflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "name"}, {Name: "spell"}},
	}
	if ignoreDup {
		onConflict.DoNothing = true
	} else {
		onConflict.DoUpdates = clause.AssignmentColumns(updateCols)
	}
	clauses = append(clauses, onConflict)
	// 执行 upsert 并获取结果
	tx := configs.DB.Model(&model.StationDO{}).WithContext(ctx).Clauses(clauses...).Create(&stations)
	return tx.RowsAffected, tx.Error
}
