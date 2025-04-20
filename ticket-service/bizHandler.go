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
	return service.NewTicketService().PageListTicketQuery(ctx, request)
}

func (t TicketServiceHandler) PurchaseTickets(ctx context.Context, request *ticket_service.PurchaseTicketRequest) (*ticket_service.PurchaseTicketResponse, error) {
	return service.NewTicketService().PurchaseTickets(ctx, request)
}

func (t TicketServiceHandler) CancelTicket(ctx context.Context, request *ticket_service.CancelTicketRequest) (*ticket_service.CancelTicketResponse, error) {
	return service.NewTicketService().CancelTicketOrder(ctx, request)
}

func (t TicketServiceHandler) GetTicket(ctx context.Context, request *ticket_service.GetTicketRequest) (*ticket_service.GetTicketResponse, error) {
	return service.NewTicketService().GetTicket(ctx, request)
}
