package configs

import (
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"google.golang.org/grpc"
	"os"
)

var UserServiceClient user_service.UserServiceClient

// InitUserServiceClient init rpc client
func InitUserServiceClient() {
	host := os.Getenv("NGINX_HOST")
	if host == "" {
		host = "localhost"
	}
	conn, err := grpc.NewClient(host+":8849", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	UserServiceClient = user_service.NewUserServiceClient(conn)
}
