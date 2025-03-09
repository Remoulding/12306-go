package service

import (
	"context"
	"fmt"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/biz/model"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/bytedance/sonic"
	"strconv"
)

var log = configs.Log

type RegionStationService struct {
}

func NewRegionStationService() *RegionStationService {
	return &RegionStationService{}
}

var regionStationQueryTypeEnum = map[int32][]string{
	1: {"A", "B", "C", "D", "E"},
	2: {"F", "G", "H", "R", "J"},
	3: {"K", "L", "M", "N", "O"},
	4: {"P", "Q", "R", "S", "T"},
	5: {"U", "V", "W", "X", "Y", "Z"},
}

func (s *RegionStationService) ListRegionStation(ctx context.Context, req *ticket_service.ListRegionStationRequest) (*ticket_service.ListRegionStationResponse, error) {
	resp := &ticket_service.ListRegionStationResponse{}
	name := req.GetName()
	if len(name) > 0 {
		data, err := SafeLoad(ctx, configs.RegionStation+name, fmt.Sprintf(configs.LockQueryRegionStationList, name),
			func(ctx context.Context) (interface{}, error) {
				return dao.QueryStationByName(ctx, "test")
			})
		if err != nil {
			resp.Message = err.Error()
		}
		var station model.StationDO
		err = sonic.UnmarshalString(data, &station)
		if err != nil {
			resp.Message = "查询车站失败"
			return resp, nil
		}
		resp.Success = true
		resp.Data = append(resp.Data, &ticket_service.RegionStation{
			Name:  station.Name,
			Code:  station.Code,
			Spell: station.Spell,
		})
		return resp, nil
	}
	condition := map[string]interface{}{}
	// query region
	queryType := req.GetQueryType()
	if queryType == 0 {
		condition["popular_flag = ?"] = 1
	} else {
		typeEnum := regionStationQueryTypeEnum[queryType]
		if typeEnum == nil {
			resp.Message = "查询类型错误"
			return resp, nil
		}
		condition["initial in ?"] = typeEnum
	}
	queryTypeStr := strconv.Itoa(int(queryType))

	data, err := SafeLoad(ctx, configs.RegionStation+queryTypeStr, fmt.Sprintf(configs.LockQueryRegionStationList, queryTypeStr),
		func(ctx context.Context) (interface{}, error) {
			return dao.QueryRegions(ctx, condition)
		})
	if err != nil {
		resp.Message = err.Error()
		return resp, nil
	}
	var regions []*model.RegionDO
	err = sonic.UnmarshalString(data, &regions)
	if err != nil {
		resp.Message = "查询地区失败"
		log.WithContext(ctx).Errorf("unmarshal region failed, err: %v", err)
		return resp, nil
	}
	resp.Success = true
	for _, region := range regions {
		resp.Data = append(resp.Data, &ticket_service.RegionStation{
			Name:  region.Name,
			Code:  region.Code,
			Spell: region.Spell,
		})
	}
	return resp, nil
}

func (s *RegionStationService) ListAllRegionStation(ctx context.Context) (*ticket_service.ListAllStationResponse, error) {
	resp := &ticket_service.ListAllStationResponse{}
	data, err := SafeLoad(ctx, configs.StationAll, configs.LockQueryAllRegionStationList, func(ctx context.Context) (interface{}, error) {
		return dao.QueryStations(ctx, nil)
	})
	if err != nil {
		resp.Message = "查询地区失败"
		log.WithContext(ctx).Errorf("safeLoad station failed, err: %v", err)
		return resp, nil
	}
	var stations []*model.StationDO
	err = sonic.UnmarshalString(data, &stations)
	if err != nil {
		resp.Message = "查询地区失败"
		log.WithContext(ctx).Errorf("unmarshal station failed, err: %v", err)
		return resp, nil
	}
	for _, station := range stations {
		resp.Data = append(resp.Data, &ticket_service.Station{
			Name:       station.Name,
			Code:       station.Code,
			Spell:      station.Spell,
			RegionName: station.RegionName,
		})
	}
	resp.Success = true
	return resp, nil
}
