package main

import (
	"context"
	"errors"
	"github.com/Remoulding/12306-go/idl-gen/user_service"
	"github.com/Remoulding/12306-go/user-service/configs"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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

func InitWebServer(ctx context.Context) *http.Server {
	// 等待 gRPC 服务启动
	time.Sleep(1 * time.Second)
	rpcPort := os.Getenv("GRPC_PORT")
	httpPort := os.Getenv("HTTP_PORT")
	if rpcPort == "" {
		rpcPort = "50050"
	}
	if httpPort == "" {
		httpPort = "8080"
	}
	// 连接 gRPC 服务器
	conn, err := grpc.DialContext(ctx, "localhost:"+rpcPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}

	gwMux := runtime.NewServeMux(
		runtime.WithMetadata(customHeaderMatcher),
	)
	err = user_service.RegisterUserServiceHandler(ctx, gwMux, conn)
	if err != nil {
		log.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	gwServer := &http.Server{
		Addr:    "0.0.0.0:" + httpPort,
		Handler: gwMux,
	}

	log.Println("🚀 gRPC-Gateway server is running at port", httpPort)

	go func() {
		if err = gwServer.ListenAndServe(); err != nil && !errors.Is(http.ErrServerClosed, err) {
			log.Fatalf("gRPC-Gateway Server failed: %v", err)
		}
	}()

	// 监听 context 退出信号
	go func() {
		<-ctx.Done()
		log.Println("⏳ Shutting down gRPC-Gateway server...")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = gwServer.Shutdown(shutdownCtx); err != nil {
			log.Fatalf("Failed to shutdown gRPC-Gateway: %v", err)
		}
	}()

	return gwServer
}
