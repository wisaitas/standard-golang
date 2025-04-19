package user

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type Transaction interface {
}

type transaction struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewTransaction(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Transaction {
	return &transaction{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
