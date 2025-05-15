package main

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configs.InitLog()
	configs.InitDBInstance()
	configs.InitCache()
	configs.InitRedSync()
	// 初始化client
	configs.InitUserServiceClient()
	// 创建全局 context，用于优雅退出
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 gRPC 服务
	InitRpcServer(ctx)

	// 启动 gRPC Gateway
	InitWebServer(ctx)

	// 启动 Canal 服务
	//InitCanal(ctx)

	// 监听系统信号，优雅退出
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan
	log.Println("🚦 Received shutdown signal")

	// 触发 context 取消
	cancel()

	// 额外等待 3s 让 goroutine 执行清理
	time.Sleep(3 * time.Second)
	log.Println("✅ Server shutdown complete")
}
