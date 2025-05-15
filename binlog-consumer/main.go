package main

import (
	"context"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/go-mysql-org/go-mysql/canal"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	mysqlAddr     = "127.0.0.1:3306"
	mysqlUser     = "canal"
	mysqlPassword = "canal_password"
)

func main() {
	configs.InitDBInstance()
	configs.InitCache()
	configs.InitRedSync()
	// 创建全局 context，用于优雅退出
	ctx, cancel := context.WithCancel(context.Background())

	// 启动 Canal 服务
	InitCanal(ctx)

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

// InitCanal 初始化Canal客户端
func InitCanal(ctx context.Context) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = mysqlAddr
	cfg.User = mysqlUser
	cfg.Password = mysqlPassword
	cfg.Dump.ExecutionPath = "" // 不需要mysqldump
	cfg.Dump.DiscardErr = false
	// 只订阅特定库表
	cfg.IncludeTableRegex = []string{"train_db\\.t_seat"}
	cfg.ServerID = 1235 // 随便设置一个不重复的
	cfg.Charset = "utf8mb4"
	cfg.Flavor = "mysql"

	c, err := canal.NewCanal(cfg)
	if err != nil {
		log.Fatalf("创建Canal客户端失败: %v", err)
	}
	// 设置事件处理器
	handler := NewSeatEventHandler(configs.Cache)
	c.SetEventHandler(handler)
	// 启动Canal客户端
	go func() {
		if err = c.RunFrom(handler.LoadPosition()); err != nil {
			log.Fatalf("Canal客户端运行失败: %v", err)
		}
	}()
	//go func() {
	//	if err = c.Run(); err != nil {
	//		log.Fatalf("Canal客户端运行失败: %v", err)
	//	}
	//}()

	// 监听 context 退出信号
	go func() {
		<-ctx.Done()
		log.Println("⏳ Shutting down canal client...")
		c.Close()
	}()

}
