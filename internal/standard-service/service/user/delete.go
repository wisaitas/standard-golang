package user

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"github.com/wisaitas/standard-golang/pkg"
)

type Delete interface {
}

type delete struct {
	userRepository repository.UserRepository
	redisUtil      pkg.Redis
}

func NewDelete(
	userRepository repository.UserRepository,
	redisUtil pkg.Redis,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
