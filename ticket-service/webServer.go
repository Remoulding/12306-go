package main

import (
	"context"
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"strings"
)

// 注入 HTTP Header 到 context
func customHeaderMatcher(ctx context.Context, req *http.Request) metadata.MD {
	md := make(metadata.MD)
	for key, values := range req.Header {
		for _, value := range values {
			if key == configs.UserIdKey || key == configs.UserNameKey ||
				key == configs.RealNameKey || key == configs.UserTokenKey {
				log.Println(key, value)
				md.Append(strings.ToLower(key), value)
			}
		}

	}
	return md
}

func InitWebServer() {
	// 初始化 grpc gateway
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close connection: %v", err)
		}
	}()
	gwMux := runtime.NewServeMux(
		runtime.WithMetadata(customHeaderMatcher),
	)
	err = user_service.RegisterUserServiceHandler(context.Background(), gwMux, conn)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}
	// http服务，使用rpc客户端
	gwServer := &http.Server{
		Addr:    ":8082",
		Handler: gwMux,
	}
	log.Println("gRPC-Gateway server is running at port 50052")
	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
