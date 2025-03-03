package initial

import (
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/configs"
	"gorm.io/gorm"
)

type Configs struct {
	DB    *gorm.DB
	Redis *redis.Client
}

func initializeConfigs() *Configs {
	configs.LoadEnv()

	db := configs.ConnectDB()

	redis := configs.ConnectRedis()

	return &Configs{
		DB:    db,
		Redis: redis,
	}
}
