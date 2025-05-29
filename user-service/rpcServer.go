package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func InitRpcServer(ctx context.Context) (*grpc.Server, net.Listener) {
	port := os.Getenv("RPC_PORT")
	if port == "" {
		port = "50050"
	}
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user_service.RegisterUserServiceServer(s, NewUserServiceHandler())
	log.Println("🚀 gRPC server is running at port ", port)

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
