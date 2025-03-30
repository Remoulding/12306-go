package configs

import "github.com/go-redis/redis/v8"

var Cache *redis.Client

func InitCache() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if _, err := Cache.Ping(Cache.Context()).Result(); err != nil {
		Log.Fatalf("🔴 Redis 连接失败: %v", err)
	}
}
