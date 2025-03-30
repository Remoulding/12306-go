package dao

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
)

func QueryRegions(ctx context.Context, condition map[string]interface{}) ([]*model.RegionDO, error) {
	var regionDO []*model.RegionDO
	query := configs.DB.Model(&model.RegionDO{})
	for exp, val := range condition {
		query = query.Where(exp, val)
	}
	err := query.Scan(&regionDO).Error
	if err != nil {
		configs.Log.WithContext(ctx).Errorf("query region failed, err: %v", err)
		return nil, err
	}
	return regionDO, nil
}
