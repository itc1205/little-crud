package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/itc1205/little-crud/internal/repository"
)

type RedisCache struct {
	r   *redis.Client
	ttl time.Duration
}

func (rc *RedisCache) New(r *redis.Client, ttl time.Duration) *RedisCache {
	return &RedisCache{r: r, ttl: ttl}
}

// Gets cache entry
func (rc *RedisCache) Get(ctx context.Context, id int32) (*repository.Goods, error) {
	goods := new(repository.Goods)
	result, err := rc.r.Get(ctx, fmt.Sprint(id)).Result()

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(result), &goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}

// Creates new cache entry
func (rc *RedisCache) Create(ctx context.Context, goods *repository.Goods) error {
	goods_json, err := json.Marshal(goods)
	if err != nil {
		return err
	}
	_, err = rc.r.Set(ctx, fmt.Sprint(goods.ID), goods_json, rc.ttl).Result()
	return err
}

// Invalidates whole cache
func (rc *RedisCache) Invalidate(ctx context.Context) error {
	_, err := rc.r.FlushDB(ctx).Result()
	return err
}
