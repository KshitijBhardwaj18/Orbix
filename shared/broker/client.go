package broker

import (
	"os"
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	rdb	*redis.Client
	ctx context.Context
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr: getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB: 0,
	})

	return &RedisClient{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *RedisClient) Ping() error {
	return r.rdb.Ping(r.ctx).Err()
}

func (r *RedisClient) Close() error {
	return r.rdb.Close()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != ""{
		return value
	}
	
	return defaultValue
}