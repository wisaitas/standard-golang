package initial

import (
	"github.com/wisaitas/standard-golang/internal/services"
	"github.com/wisaitas/standard-golang/internal/utils"
)

func initializeServices(repos *Repositories, redisClient utils.RedisClient) *Services {
	return &Services{
		UserService: services.NewUserService(repos.UserRepository, redisClient),
		AuthService: services.NewAuthService(repos.UserRepository, redisClient),
	}
}

type Services struct {
	UserService services.UserService
	AuthService services.AuthService
}
