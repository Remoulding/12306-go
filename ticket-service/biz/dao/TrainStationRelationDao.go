package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
)

func QueryTrainStationRelation(ctx context.Context, condition map[string]interface{}) ([]*model.TrainStationRelationDO, error) {
	// 查询车站关系
	var trainStationRelations []*model.TrainStationRelationDO
	query := db.Model(&model.TrainStationRelationDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	err := query.Scan(&trainStationRelations).Error
	if err != nil {
		log.WithContext(ctx).Errorf("query train station relation failed, err: %v", err)
		return nil, err
	}
	return trainStationRelations, nil
}
