package redisconn

import (
	"project/config"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func ReddisConc() *redis.Client {
	cfg := config.GetConfig()
	db, _ := strconv.Atoi(cfg.RedisConfig.DB)
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConfig.Adress,
		Password: cfg.RedisConfig.Password,
		DB:       db,
	})
	return client
}
