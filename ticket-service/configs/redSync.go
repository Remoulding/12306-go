package configs

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

var Rs *redsync.Redsync

func InitRedSync() {
	// 创建redsync的客户端连接池
	pool := goredis.NewPool(Cache)

	// 创建redsync实例
	Rs = redsync.New(pool)
}
