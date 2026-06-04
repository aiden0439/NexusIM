package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/aiden0439/NexusIM/pkg/config"
)

func New(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
}
