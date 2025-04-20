package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func InitRpcServer(ctx context.Context) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//grpc.UnaryInterceptor(loggingUnaryInterceptor)
	ticket_service.RegisterTicketServiceServer(s, NewTicketServiceHandler())
	log.Println("🚀 gRPC server is running at port 50052")

	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("gRPC Server failed: %v", err)
		}
	}()

	// 监听 context 退出信号
	go func() {
		<-ctx.Done()
		log.Println("⏳ Shutting down gRPC server...")
		s.GracefulStop()
		err = listen.Close()
		if err != nil {
			return
		}
	}()

	return s, listen
}

// Unary 拦截器，打印请求和响应
func loggingUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	// 打印请求信息
	configs.Log.WithContext(ctx).Infof("[gRPC Request] Method: %s | Request: %+v", info.FullMethod, req)

	// 执行 gRPC 方法
	resp, err := handler(ctx, req)

	// 打印响应信息
	configs.Log.WithContext(ctx).Infof("[gRPC Response] Method: %s | Duration: %v | Response: %+v | Error: %v",
		info.FullMethod, time.Since(start), resp, err)
	return resp, err
}
