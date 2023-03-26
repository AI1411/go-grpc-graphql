package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/AI1411/go-grpc-graphql/internal/env"
)

type RedisClient struct {
	redisClient *redis.Client
}

func NewRedisClient(e *env.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", e.RedisHost, e.RedisPort),
	})
}
