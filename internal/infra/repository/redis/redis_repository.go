package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) error
	Delete(ctx context.Context, key string) error
}

type redisRepository struct {
	redisClient *redis.Client
}

func NewRedisRepository(redisClient *redis.Client) RedisRepository {
	return &redisRepository{
		redisClient: redisClient,
	}
}

func (r *redisRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (r *redisRepository) Set(ctx context.Context, key, value string) error {
	if err := r.redisClient.Set(ctx, key, value, time.Hour*24).Err(); err != nil {
		return err
	}

	return nil
}

func (r *redisRepository) Delete(ctx context.Context, key string) error {
	if err := r.redisClient.Del(ctx, key).Err(); err != nil {
		return err
	}

	return nil
}
