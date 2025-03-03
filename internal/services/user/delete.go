package user

import (
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type Delete interface {
}

type delete struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewDelete(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
