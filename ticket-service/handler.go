package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/ticket-service/biz/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TicketServiceHandler struct {
	ticket_service.UnimplementedTicketServiceServer
}

func NewTicketServiceHandler() ticket_service.TicketServiceServer {
	return &TicketServiceHandler{}
}

func (t TicketServiceHandler) ListRegionStation(ctx context.Context, request *ticket_service.ListRegionStationRequest) (*ticket_service.ListRegionStationResponse, error) {
	return service.NewRegionStationService().ListRegionStation(ctx, request)
}

func (t TicketServiceHandler) ListAllStation(ctx context.Context, empty *emptypb.Empty) (*ticket_service.ListAllStationResponse, error) {
	return service.NewRegionStationService().ListAllRegionStation(ctx)
}

func (t TicketServiceHandler) ListTrainStationQuery(ctx context.Context, request *ticket_service.TrainStationQueryRequest) (*ticket_service.TrainStationQueryResponse, error) {
	return service.NewTrainStationService().ListTrainStationQuery(ctx, request)
}

func (t TicketServiceHandler) PageListTicketQuery(ctx context.Context, request *ticket_service.TicketPageQueryRequest) (*ticket_service.TicketPageQueryResponse, error) {
	//TODO implement me
	// 查一下就完事
	panic("implement me")
}

func (t TicketServiceHandler) PurchaseTickets(ctx context.Context, request *ticket_service.PurchaseTicketRequest) (*ticket_service.PurchaseTicketResponse, error) {
	//TODO implement me
	// 买票锁全车，孩子你无敌了，搞个策略模式应付一下得了
	panic("implement me")
}

func (t TicketServiceHandler) CancelTicketOrder(ctx context.Context, request *ticket_service.CancelTicketOrderRequest) (*ticket_service.CancelTicketOrderResponse, error) {
	//TODO implement me
	// 回滚一下就完事
	panic("implement me")
}

func (t TicketServiceHandler) GetPayInfo(ctx context.Context, request *ticket_service.PayInfoRequest) (*ticket_service.PayInfoResponse, error) {
	//TODO implement me
	// 老子不想写了
	return nil, nil
}
