package user

import (
	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type Delete interface {
}

type delete struct {
	userRepository repository.UserRepository
	redisUtil      redisPkg.Redis
}

func NewDelete(
	userRepository repository.UserRepository,
	redisUtil redisPkg.Redis,
) Delete {
	return &delete{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
