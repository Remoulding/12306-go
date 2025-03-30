package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
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
