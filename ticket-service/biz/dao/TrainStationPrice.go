package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

func QueryTrainStationPrice(ctx context.Context, condition map[string]interface{}) ([]*model.TrainStationPriceDO, error) {
	// 查询车站票价
	var trainStationPrices []*model.TrainStationPriceDO
	err := configs.DB.WithContext(ctx).Model(&model.TrainStationPriceDO{}).Where(condition).Scan(&trainStationPrices).Error
	return trainStationPrices, err
}
