package routes

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
	"github.com/wisaitas/standard-golang/internal/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	app          fiber.Router
	userHandler  *handlers.UserHandler
	userValidate *validates.UserValidate
}

func NewUserRoutes(
	app fiber.Router,
	userHandler *handlers.UserHandler,
	userValidate *validates.UserValidate,
) *UserRoutes {
	return &UserRoutes{
		app:          app,
		userHandler:  userHandler,
		userValidate: userValidate,
	}
}

func (r *UserRoutes) UserRoutes() {
	users := r.app.Group("/users")
	users.Get("/", r.userHandler.GetUsers)
	users.Post("/", r.userValidate.ValidateCreateUserRequest, r.userHandler.CreateUser)
}
