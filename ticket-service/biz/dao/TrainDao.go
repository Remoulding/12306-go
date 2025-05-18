package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"gorm.io/gorm/clause"
)

func QueryTrainById(ctx context.Context, trainID int64) (*model.TrainDO, error) {
	train := &model.TrainDO{}
	err := configs.DB.Model(&model.TrainDO{}).Where("id = ?", trainID).First(train).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query train failed, err: %v", err)
		return nil, err
	}
	return train, nil
}

func QueryTrain(ctx context.Context, condition map[string]interface{}) ([]*model.TrainDO, error) {
	var trains []*model.TrainDO
	query := configs.DB.Model(&model.TrainDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	if err := query.Scan(&trains).Error; err != nil {
		configs.Log.WithContext(ctx).Errorf("query train failed, err: %v", err)
		return nil, err
	}
	return trains, nil
}

func UpsertTrains(ctx context.Context, trains []*model.TrainDO, ignoreDup bool, updateCols []string) (int64, error) {
	// 使用clause
	var clauses []clause.Expression
	onConflict := clause.OnConflict{
		Columns: []clause.Column{{Name: "train_no"}, {Name: "train_type"}},
	}
	if ignoreDup {
		onConflict.DoNothing = true
	} else {
		onConflict.DoUpdates = clause.AssignmentColumns(updateCols)
	}
	clauses = append(clauses, onConflict)
	// 执行 upsert 并获取结果
	tx := configs.DB.Model(&model.TrainDO{}).WithContext(ctx).Clauses(clauses...).Create(&trains)
	return tx.RowsAffected, tx.Error
}
