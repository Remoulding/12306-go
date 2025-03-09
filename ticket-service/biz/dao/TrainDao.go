package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
)

func QueryTrainById(ctx context.Context, trainID int64) (*model.TrainDO, error) {
	var train *model.TrainDO
	err := db.Model(&model.TrainDO{}).Where("train_id = ?", trainID).First(train).Error
	if err != nil {
		log.WithContext(ctx).Errorf("query train failed, err: %v", err)
		return nil, err
	}
	return train, nil
}

func QueryTrain(ctx context.Context, condition map[string]interface{}) ([]*model.TrainDO, error) {
	var trains []*model.TrainDO
	query := db.Model(&model.TrainDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	if err := query.Scan(&trains).Error; err != nil {
		log.WithContext(ctx).Errorf("query train failed, err: %v", err)
		return nil, err
	}
	return trains, nil
}
