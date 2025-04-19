package user

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type Delete interface {
}

type delete struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewDelete(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
