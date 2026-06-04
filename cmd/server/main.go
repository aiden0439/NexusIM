package main

import (
	"context"
	"log"

	"github.com/zhf0439/im-server/internal/router"
	"github.com/zhf0439/im-server/pkg/config"
	"github.com/zhf0439/im-server/pkg/kafka"
	"github.com/zhf0439/im-server/pkg/logger"
	"github.com/zhf0439/im-server/pkg/mysql"
	"github.com/zhf0439/im-server/pkg/redis"
	"go.uber.org/zap"
)

func main() {
	cfg, err := config.Load("config/config.yaml")
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	if err := logger.Init(cfg.Log); err != nil {
		log.Fatalf("init logger: %v", err)
	}
	defer logger.Sync()

	db, err := mysql.New(cfg.MySQL)
	if err != nil {
		logger.L().Fatal("connect mysql", zap.Error(err))
	}

	redisClient := redis.New(cfg.Redis)
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		logger.L().Fatal("connect redis", zap.Error(err))
	}

	kafkaClient := kafka.New(cfg.Kafka)
	defer kafkaClient.Close()

	r := router.New(router.Dependencies{
		Config: cfg,
		DB:     db,
		Redis:  redisClient,
		Kafka:  kafkaClient,
	})

	logger.L().Info("server starting", zap.String("addr", cfg.Server.Addr))
	if err := r.Run(cfg.Server.Addr); err != nil {
		logger.L().Fatal("server stopped", zap.Error(err))
	}
}
