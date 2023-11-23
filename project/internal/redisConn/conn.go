package redisconn

import (
	"github.com/redis/go-redis/v9"
)

func ReddisConc() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
	return client
}
