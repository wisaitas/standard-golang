package initial

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
)

func initializeHandlers(services *Services) *Handlers {
	return &Handlers{
		UserHandler: *handlers.NewUserHandler(services.UserService),
		AuthHandler: *handlers.NewAuthHandler(services.AuthService),
	}
}

type Handlers struct {
	UserHandler handlers.UserHandler
	AuthHandler handlers.AuthHandler
}
