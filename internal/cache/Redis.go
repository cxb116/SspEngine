package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/yourusername/ssp_grpc/internal/config"
)

func NewRedisClient(cfg *config.Redis) (*redis.Client, error) {
	opts := &redis.Options{
		Addr:     cfg.Addr,
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.Db,
	}
	redisClient := redis.NewClient(opts)
	if status := redisClient.Ping(context.Background()); status.Err() != nil {
		return nil, status.Err()
	}
	return redisClient, nil
}
