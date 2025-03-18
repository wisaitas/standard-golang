package user

import (
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type Transaction interface {
}

type transaction struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisClient
}

func NewTransaction(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisClient,
) Transaction {
	return &transaction{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
