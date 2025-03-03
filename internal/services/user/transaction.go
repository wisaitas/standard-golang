package user

import (
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type Transaction interface {
}

type transaction struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewTransaction(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) Transaction {
	return &transaction{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}
