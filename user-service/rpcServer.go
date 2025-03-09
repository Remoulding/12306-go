package main

import (
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitRpcServer() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	user_service.RegisterUserServiceServer(s, NewUserServiceHandler())
	log.Println("gRPC server is running at port 50051")
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
