package rpc

import (
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"google.golang.org/grpc"
)

var UserServiceClient user_service.UserServiceClient

// InitClient init rpc client
func InitClient() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	UserServiceClient = user_service.NewUserServiceClient(conn)
}
