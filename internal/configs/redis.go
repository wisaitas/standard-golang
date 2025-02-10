package configs

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ENV.REDIS_HOST, ENV.REDIS_PORT),
		Password: "",
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("error connecting to redis: %v\n", err)
	}

	log.Println("redis connected successfully")

	return redisClient
}
