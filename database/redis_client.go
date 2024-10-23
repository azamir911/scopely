package database

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

func NewRedisClient(ctx context.Context) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis server address
		Password: "",               // No password set for Redis by default
		DB:       0,                // Use default DB
	})

	// Test connection
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	return client
}
