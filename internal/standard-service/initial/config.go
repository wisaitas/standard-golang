package initial

import (
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/config"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"gorm.io/gorm"
)

type clientConfig struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func newClientConfig() *clientConfig {
	db := config.ConnectDatabaseSQL(
		env.DB_HOST,
		env.DB_PORT,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
		env.DB_DRIVER,
	)

	redis := config.ConnectRedis(
		env.REDIS_HOST,
		env.REDIS_PORT,
		env.REDIS_PASSWORD,
	)

	return &clientConfig{
		DB:    db,
		Redis: redis,
	}
}
