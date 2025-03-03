package user

import (
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type Update interface {
}

type update struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewUpdate(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) Update {
	return &update{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
