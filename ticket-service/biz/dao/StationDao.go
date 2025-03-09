package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

var db = configs.DB
var log = configs.Log

func QueryStationByName(ctx context.Context, name string) (*model.StationDO, error) {
	var station *model.StationDO
	name += "%"
	err := db.Model(&model.StationDO{}).Where("name like ?", name).
		Or("spell like ?", name).First(station).Error
	if err != nil {
		log.WithContext(ctx).Errorf("query station failed, err: %v", err)
		return nil, err
	}
	return station, nil
}

func QueryStations(ctx context.Context, condition map[string]interface{}) ([]*model.StationDO, error) {
	var stations []*model.StationDO
	query := db.Model(&model.StationDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	if err := query.Scan(&stations).Error; err != nil {
		log.WithContext(ctx).Errorf("query station failed, err: %v", err)
		return nil, err
	}
	return stations, nil
}
