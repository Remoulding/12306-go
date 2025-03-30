package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

func QueryTrainStationRelation(ctx context.Context, condition map[string]interface{}) ([]*model.TrainStationRelationDO, error) {
	// 查询车站关系
	var trainStationRelations []*model.TrainStationRelationDO
	query := configs.DB.Model(&model.TrainStationRelationDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	err := query.Scan(&trainStationRelations).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query train station relation failed, err: %v", err)
		return nil, err
	}
	return trainStationRelations, nil
}
