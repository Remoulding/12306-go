package configs

import "github.com/go-redis/redis/v8"

var Cache *redis.Client

func InitCache() {
	Cache = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
