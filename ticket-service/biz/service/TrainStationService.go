package service

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/ticket-service/biz/dao"
	"github.com/Remoulding/12306-go/ticket-service/tools"
)

type TrainStationService struct {
}

func NewTrainStationService() *TrainStationService {
	return &TrainStationService{}
}

type Route struct {
	StartStation string
	EndStation   string
}

func (s *TrainStationService) ListTrainStationQuery(ctx context.Context, request *ticket_service.TrainStationQueryRequest) (*ticket_service.TrainStationQueryResponse, error) {
	resp := &ticket_service.TrainStationQueryResponse{}
	condition := map[string]interface{}{
		"train_id = ?": request.GetTrainId(),
	}
	trainStations, err := dao.QueryTrainStation(ctx, condition)
	if err != nil {
		resp.Message = "查询车站失败"
		log.WithContext(ctx).Errorf("query train station failed, err: %v", err)
		return resp, nil
	}
	data := make([]*ticket_service.TrainStation, 0, len(trainStations))
	for _, station := range trainStations {
		data = append(data, &ticket_service.TrainStation{
			Sequence:      station.Sequence,
			Departure:     station.Departure,
			ArrivalTime:   tools.ConvertTimeToStr(station.ArrivalTime),
			DepartureTime: tools.ConvertTimeToStr(station.DepartureTime),
			StopoverTime:  int32(station.StopoverTime),
		})
	}
	resp.Data = data
	resp.Success = true
	return resp, nil
}

// ListTrainStationRoute 获取开始站点和目的站点及中间站点信息
func (s *TrainStationService) ListTrainStationRoute(ctx context.Context, trainId, departure, arrival string) ([]*Route, error) {
	condition := map[string]interface{}{
		"train_id = ?": trainId,
	}
	trainStations, err := dao.QueryTrainStation(ctx, condition)
	if err != nil {
		log.WithContext(ctx).Errorf("query train station failed, err: %v", err)
		return nil, err
	}
	var stations []string
	for _, station := range trainStations {
		stations = append(stations, station.Departure)
	}
	return s.throughStation(stations, departure, arrival), nil
}

// ListTakeoutTrainStationRoute 获取开始站点和目的站点、中间站点以及关联站点信息
func (s *TrainStationService) ListTakeoutTrainStationRoute(ctx context.Context, trainId, departure, arrival string) ([]*Route, error) {
	condition := map[string]interface{}{
		"train_id = ?": trainId,
	}
	trainStations, err := dao.QueryTrainStation(ctx, condition)
	if err != nil {
		log.WithContext(ctx).Errorf("query train station failed, err: %v", err)
		return nil, err
	}
	var stations []string
	for _, station := range trainStations {
		stations = append(stations, station.Departure)
	}
	return s.takeoutStation(stations, departure, arrival), nil
}

func (s *TrainStationService) throughStation(stations []string, startStation, endStation string) []*Route {
	var routesToDeduct []*Route
	startIndex := -1
	endIndex := -1

	// 获取出发站和终点站的索引
	for i, station := range stations {
		if station == startStation {
			startIndex = i
		}
		if station == endStation {
			endIndex = i
		}
	}

	if startIndex == -1 || endIndex == -1 || startIndex >= endIndex {
		return routesToDeduct
	}

	// 生成路径
	for i := startIndex; i < endIndex; i++ {
		for j := i + 1; j <= endIndex; j++ {
			routesToDeduct = append(routesToDeduct, &Route{stations[i], stations[j]})
		}
	}
	return routesToDeduct
}

func (s *TrainStationService) takeoutStation(stations []string, startStation, endStation string) []*Route {
	var takeoutStationList []*Route
	startIndex := -1
	endIndex := -1

	// 获取出发站和终点站的索引
	for i, station := range stations {
		if station == startStation {
			startIndex = i
		}
		if station == endStation {
			endIndex = i
		}
	}

	if startIndex == -1 || endIndex == -1 || startIndex >= endIndex {
		return takeoutStationList
	}

	// 排除出发站之前的站点
	if startIndex != 0 {
		for i := 0; i < startIndex; i++ {
			for j := 1; j < len(stations)-startIndex; j++ {
				takeoutStationList = append(takeoutStationList, &Route{stations[i], stations[startIndex+j]})
			}
		}
	}

	// 生成路径
	for i := startIndex; i <= endIndex; i++ {
		for j := i + 1; j < len(stations) && i < endIndex; j++ {
			takeoutStationList = append(takeoutStationList, &Route{stations[i], stations[j]})
		}
	}
	return takeoutStationList
}
