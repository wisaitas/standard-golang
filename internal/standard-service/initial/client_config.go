package initial

import (
	"context"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	standardservice "github.com/wisaitas/standard-golang/internal/standard-service"
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
		standardservice.ENV.Database.Host,
		standardservice.ENV.Database.Port,
		standardservice.ENV.Database.User,
		standardservice.ENV.Database.Password,
		standardservice.ENV.Database.Name,
	)

	redis := ConnectRedis(
		standardservice.ENV.Redis.Host,
		standardservice.ENV.Redis.Port,
		standardservice.ENV.Redis.Password,
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

	log.Println("database connected successfully")

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
