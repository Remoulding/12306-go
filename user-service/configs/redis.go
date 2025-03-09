package configs

import "github.com/go-redis/redis/v8"

var UserLoginCache *redis.Client

func InitCache() {
	UserLoginCache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
