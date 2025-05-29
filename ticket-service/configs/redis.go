package configs

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var Cache *redis.Client

func InitCache() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	Cache = redis.NewClient(&redis.Options{
		Addr: host + ":6379",
	})
	if _, err := Cache.Ping(Cache.Context()).Result(); err != nil {
		Log.Fatalf("🔴 Redis 连接失败: %v", err)
	}
}
