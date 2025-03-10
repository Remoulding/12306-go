package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitRpcServer(ctx context.Context) (*grpc.Server, net.Listener) {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
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
