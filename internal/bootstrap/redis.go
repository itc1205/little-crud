package bootstrap

import (
	"context"
	"fmt"

	"github.com/itc1205/little-crud/internal/config"
	"github.com/redis/go-redis/v9"
)

func InitRedisClient(ctx context.Context, cfg config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(redisOptionsFromConfig(cfg))
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}

func redisOptionsFromConfig(cfg config.RedisConfig) *redis.Options {
	addr := fmt.Sprintf("%s:%s", cfg.RHost, cfg.RPort)
	return &redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
}
