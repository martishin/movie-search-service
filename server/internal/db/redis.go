package db

import (
	"context"
	"fmt"
	"time"

	"github.com/martishin/movie-search-service/internal/model/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisClient(config *config.RedisConfig) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", config.Host, config.Port),
		DB:   config.DB,
	})

	// Test Redis connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := client.Ping(ctx).Result()

	if err != nil {
		return client, err
	}

	return client, nil
}
