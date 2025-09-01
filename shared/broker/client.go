package broker

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

type Broker struct {
	rdb *redis.Client
	ctx context.Context
}

func NewRedisClient() *Broker {
	rdb := redis.NewClient(&redis.Options{
		Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       0,
	})

	return &Broker{
		rdb: rdb,
		ctx: context.Background(),
	}
}

func (r *Broker) Ping() error {
	return r.rdb.Ping(r.ctx).Err()
}

func (r *Broker) Close() error {
	return r.rdb.Close()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
