package initial

import (
	authService "github.com/wisaitas/standard-golang/internal/services/auth"
	userService "github.com/wisaitas/standard-golang/internal/services/user"
	"github.com/wisaitas/standard-golang/internal/utils"
)

func initializeServices(repos *Repositories, redisClient utils.RedisClient) *Services {
	return &Services{
		UserService: userService.NewUserService(
			userService.NewRead(repos.UserRepository, redisClient),
			userService.NewCreate(repos.UserRepository, redisClient),
			userService.NewUpdate(repos.UserRepository, redisClient),
			userService.NewDelete(repos.UserRepository, redisClient),
			userService.NewTransaction(repos.UserRepository, redisClient),
		),
		AuthService: authService.NewAuthService(repos.UserRepository, redisClient),
	}
}

type Services struct {
	UserService userService.UserService
	AuthService authService.AuthService
}
