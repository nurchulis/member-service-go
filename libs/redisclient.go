package libs

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func NewClient() *redis.Client {
	// Create a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "redis-16741.c302.asia-northeast1-1.gce.cloud.redislabs.com:16741",
		Password: "CtU0u6A8Drr7ohXAUSEier6u9FOP84mr", // no password set
		DB:       0,                                  // use default DB
	})

	// Ping the Redis server to make sure it's available
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return client
}
