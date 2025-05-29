package configs

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var UserLoginCache *redis.Client

func InitCache() {
	host := os.Getenv("REDIS_HOST")
	if host == "" {
		host = "localhost"
	}
	UserLoginCache = redis.NewClient(&redis.Options{
		Addr: host + ":6379",
	})
}
