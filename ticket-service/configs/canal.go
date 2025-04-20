package configs

import (
	"fmt"
	"github.com/go-mysql-org/go-mysql/canal"
)

// 配置常量
const (
	mysqlAddr     = "127.0.0.1:3306"
	mysqlUser     = "canal_user"
	mysqlPassword = "canal_password"
)

// InitCanal 初始化Canal客户端
func InitCanal() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	cfg.Addr = mysqlAddr
	cfg.User = mysqlUser
	cfg.Password = mysqlPassword
	cfg.Dump.ExecutionPath = "" // 不需要mysqldump
	cfg.Dump.DiscardErr = false

	// 只订阅特定库表
	cfg.IncludeTableRegex = []string{"^train_db\\.t_seat$"}

	c, err := canal.NewCanal(cfg)
	if err != nil {
		return nil, fmt.Errorf("创建Canal客户端失败: %v", err)
	}

	return c, nil
}
