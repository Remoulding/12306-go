package main

import (
	"github.com/Remoulding/12306-go/idl-gen/ticket_service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitRpcServer() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	ticket_service.RegisterTicketServiceServer(s, NewTicketServiceHandler())
	log.Println("gRPC server is running at port 50052")
	go func() {
		if err := s.Serve(listen); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

}
