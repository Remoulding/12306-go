package configs

import (
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"google.golang.org/grpc"
)

var UserServiceClient user_service.UserServiceClient

// InitUserServiceClient init rpc client
func InitUserServiceClient() {
	conn, err := grpc.NewClient("localhost:8849", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	UserServiceClient = user_service.NewUserServiceClient(conn)
}
