package utils

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, keys ...string) error
	Exists(ctx context.Context, keys ...string) (bool, error)
}

type redisClient struct {
	Client *redis.Client
}

func NewRedisClient(client *redis.Client) RedisClient {
	return &redisClient{
		Client: client,
	}
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.Client.Set(ctx, key, value, expiration).Err()
}

func (r *redisClient) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}

func (r *redisClient) Del(ctx context.Context, keys ...string) error {
	return r.Client.Del(ctx, keys...).Err()
}

func (r *redisClient) Exists(ctx context.Context, keys ...string) (bool, error) {
	exists, err := r.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return false, err
	}
	return exists > 0, nil
}
