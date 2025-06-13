package initial

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type clientConfig struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func newClientConfig() *clientConfig {
	db := connectPostgresDatabase(
		env.Environment.Database.Host,
		env.Environment.Database.Port,
		env.Environment.Database.User,
		env.Environment.Database.Password,
		env.Environment.Database.Name,
	)

	redis := ConnectRedis(
		env.Environment.Redis.Host,
		env.Environment.Redis.Port,
		env.Environment.Redis.Password,
	)

	return &clientConfig{
		DB:    db,
		Redis: redis,
	}
}

func connectPostgresDatabase(
	host string,
	port int,
	user string,
	password string,
	database string,
) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Bangkok",
		host,
		user,
		password,
		database,
		port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	return db
}

func ConnectRedis(
	host string,
	port int,
	password string,
) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       0,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("error connecting to redis: %v\n", err)
	}

	log.Println("redis connected successfully")

	return redisClient
}
