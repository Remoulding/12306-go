package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
)

func QueryTrainStation(ctx context.Context, condition map[string]interface{}) ([]*model.TrainStationDO, error) {
	var trainStations []*model.TrainStationDO
	query := db.Model(&model.TrainStationDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	if err := query.Scan(&trainStations).Error; err != nil {
		log.WithContext(ctx).Errorf("query train station failed, err: %v", err)
		return nil, err
	}
	return trainStations, nil
}
