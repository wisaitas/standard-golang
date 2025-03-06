package routes

import (
	"github.com/wisaitas/standard-golang/internal/handlers"
	"github.com/wisaitas/standard-golang/internal/middlewares"
	"github.com/wisaitas/standard-golang/internal/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	app            fiber.Router
	userHandler    *handlers.UserHandler
	userValidate   *validates.UserValidate
	authMiddleware *middlewares.AuthMiddleware
	userMiddleware *middlewares.UserMiddleware
}

func NewUserRoutes(
	app fiber.Router,
	userHandler *handlers.UserHandler,
	userValidate *validates.UserValidate,
	authMiddleware *middlewares.AuthMiddleware,
	userMiddleware *middlewares.UserMiddleware,
) *UserRoutes {
	return &UserRoutes{
		app:            app,
		userHandler:    userHandler,
		userValidate:   userValidate,
		authMiddleware: authMiddleware,
		userMiddleware: userMiddleware,
	}
}

func (r *UserRoutes) UserRoutes() {
	users := r.app.Group("/users")

	// Method GET
	users.Get("/", r.userMiddleware.GetUsers, r.userValidate.ValidateGetUsersRequest, r.userHandler.GetUsers)

	// Method POST
	users.Post("/", r.userValidate.ValidateCreateUserRequest, r.userHandler.CreateUser)

	// Method PATCH
	users.Patch("/:id", r.userValidate.ValidateUpdateUserRequest, r.userHandler.UpdateUser)
}
