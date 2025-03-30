// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.1
// source: ticket/ticket_service.proto

package ticket_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TicketService_ListRegionStation_FullMethodName     = "/ticket.TicketService/ListRegionStation"
	TicketService_ListAllStation_FullMethodName        = "/ticket.TicketService/ListAllStation"
	TicketService_ListTrainStationQuery_FullMethodName = "/ticket.TicketService/ListTrainStationQuery"
	TicketService_PageListTicketQuery_FullMethodName   = "/ticket.TicketService/PageListTicketQuery"
	TicketService_PurchaseTickets_FullMethodName       = "/ticket.TicketService/PurchaseTickets"
	TicketService_CancelTicket_FullMethodName          = "/ticket.TicketService/CancelTicket"
	TicketService_GetTicket_FullMethodName             = "/ticket.TicketService/GetTicket"
)

// TicketServiceClient is the client API for TicketService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义服务
type TicketServiceClient interface {
	// 查询车站&城市站点集合信息
	ListRegionStation(ctx context.Context, in *ListRegionStationRequest, opts ...grpc.CallOption) (*ListRegionStationResponse, error)
	// 查询所有站点集合信息
	ListAllStation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAllStationResponse, error)
	// 列车站点控制层
	ListTrainStationQuery(ctx context.Context, in *TrainStationQueryRequest, opts ...grpc.CallOption) (*TrainStationQueryResponse, error)
	// 车票控制层
	PageListTicketQuery(ctx context.Context, in *TicketPageQueryRequest, opts ...grpc.CallOption) (*TicketPageQueryResponse, error)
	PurchaseTickets(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*PurchaseTicketResponse, error)
	CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error)
	GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error)
}

type ticketServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTicketServiceClient(cc grpc.ClientConnInterface) TicketServiceClient {
	return &ticketServiceClient{cc}
}

func (c *ticketServiceClient) ListRegionStation(ctx context.Context, in *ListRegionStationRequest, opts ...grpc.CallOption) (*ListRegionStationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRegionStationResponse)
	err := c.cc.Invoke(ctx, TicketService_ListRegionStation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ListAllStation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAllStationResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListAllStationResponse)
	err := c.cc.Invoke(ctx, TicketService_ListAllStation_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) ListTrainStationQuery(ctx context.Context, in *TrainStationQueryRequest, opts ...grpc.CallOption) (*TrainStationQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TrainStationQueryResponse)
	err := c.cc.Invoke(ctx, TicketService_ListTrainStationQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) PageListTicketQuery(ctx context.Context, in *TicketPageQueryRequest, opts ...grpc.CallOption) (*TicketPageQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TicketPageQueryResponse)
	err := c.cc.Invoke(ctx, TicketService_PageListTicketQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) PurchaseTickets(ctx context.Context, in *PurchaseTicketRequest, opts ...grpc.CallOption) (*PurchaseTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PurchaseTicketResponse)
	err := c.cc.Invoke(ctx, TicketService_PurchaseTickets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) CancelTicket(ctx context.Context, in *CancelTicketRequest, opts ...grpc.CallOption) (*CancelTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelTicketResponse)
	err := c.cc.Invoke(ctx, TicketService_CancelTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ticketServiceClient) GetTicket(ctx context.Context, in *GetTicketRequest, opts ...grpc.CallOption) (*GetTicketResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTicketResponse)
	err := c.cc.Invoke(ctx, TicketService_GetTicket_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TicketServiceServer is the server API for TicketService service.
// All implementations must embed UnimplementedTicketServiceServer
// for forward compatibility.
//
// 定义服务
type TicketServiceServer interface {
	// 查询车站&城市站点集合信息
	ListRegionStation(context.Context, *ListRegionStationRequest) (*ListRegionStationResponse, error)
	// 查询所有站点集合信息
	ListAllStation(context.Context, *emptypb.Empty) (*ListAllStationResponse, error)
	// 列车站点控制层
	ListTrainStationQuery(context.Context, *TrainStationQueryRequest) (*TrainStationQueryResponse, error)
	// 车票控制层
	PageListTicketQuery(context.Context, *TicketPageQueryRequest) (*TicketPageQueryResponse, error)
	PurchaseTickets(context.Context, *PurchaseTicketRequest) (*PurchaseTicketResponse, error)
	CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error)
	GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error)
	mustEmbedUnimplementedTicketServiceServer()
}

// UnimplementedTicketServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTicketServiceServer struct{}

func (UnimplementedTicketServiceServer) ListRegionStation(context.Context, *ListRegionStationRequest) (*ListRegionStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRegionStation not implemented")
}
func (UnimplementedTicketServiceServer) ListAllStation(context.Context, *emptypb.Empty) (*ListAllStationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllStation not implemented")
}
func (UnimplementedTicketServiceServer) ListTrainStationQuery(context.Context, *TrainStationQueryRequest) (*TrainStationQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTrainStationQuery not implemented")
}
func (UnimplementedTicketServiceServer) PageListTicketQuery(context.Context, *TicketPageQueryRequest) (*TicketPageQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PageListTicketQuery not implemented")
}
func (UnimplementedTicketServiceServer) PurchaseTickets(context.Context, *PurchaseTicketRequest) (*PurchaseTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PurchaseTickets not implemented")
}
func (UnimplementedTicketServiceServer) CancelTicket(context.Context, *CancelTicketRequest) (*CancelTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelTicket not implemented")
}
func (UnimplementedTicketServiceServer) GetTicket(context.Context, *GetTicketRequest) (*GetTicketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTicket not implemented")
}
func (UnimplementedTicketServiceServer) mustEmbedUnimplementedTicketServiceServer() {}
func (UnimplementedTicketServiceServer) testEmbeddedByValue()                       {}

// UnsafeTicketServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TicketServiceServer will
// result in compilation errors.
type UnsafeTicketServiceServer interface {
	mustEmbedUnimplementedTicketServiceServer()
}

func RegisterTicketServiceServer(s grpc.ServiceRegistrar, srv TicketServiceServer) {
	// If the following call pancis, it indicates UnimplementedTicketServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TicketService_ServiceDesc, srv)
}

func _TicketService_ListRegionStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRegionStationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListRegionStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListRegionStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListRegionStation(ctx, req.(*ListRegionStationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ListAllStation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListAllStation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListAllStation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListAllStation(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_ListTrainStationQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TrainStationQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).ListTrainStationQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_ListTrainStationQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).ListTrainStationQuery(ctx, req.(*TrainStationQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_PageListTicketQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TicketPageQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PageListTicketQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_PageListTicketQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PageListTicketQuery(ctx, req.(*TicketPageQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_PurchaseTickets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PurchaseTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).PurchaseTickets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_PurchaseTickets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).PurchaseTickets(ctx, req.(*PurchaseTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_CancelTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).CancelTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_CancelTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).CancelTicket(ctx, req.(*CancelTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TicketService_GetTicket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTicketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TicketServiceServer).GetTicket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TicketService_GetTicket_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TicketServiceServer).GetTicket(ctx, req.(*GetTicketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TicketService_ServiceDesc is the grpc.ServiceDesc for TicketService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TicketService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ticket.TicketService",
	HandlerType: (*TicketServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRegionStation",
			Handler:    _TicketService_ListRegionStation_Handler,
		},
		{
			MethodName: "ListAllStation",
			Handler:    _TicketService_ListAllStation_Handler,
		},
		{
			MethodName: "ListTrainStationQuery",
			Handler:    _TicketService_ListTrainStationQuery_Handler,
		},
		{
			MethodName: "PageListTicketQuery",
			Handler:    _TicketService_PageListTicketQuery_Handler,
		},
		{
			MethodName: "PurchaseTickets",
			Handler:    _TicketService_PurchaseTickets_Handler,
		},
		{
			MethodName: "CancelTicket",
			Handler:    _TicketService_CancelTicket_Handler,
		},
		{
			MethodName: "GetTicket",
			Handler:    _TicketService_GetTicket_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ticket/ticket_service.proto",
}
