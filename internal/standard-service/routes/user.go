package routes

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/handlers"
	"github.com/wisaitas/standard-golang/internal/standard-service/middlewares"
	"github.com/wisaitas/standard-golang/internal/standard-service/validates"

	"github.com/gofiber/fiber/v2"
)

type UserRoutes struct {
	app            fiber.Router
	userHandler    handlers.UserHandler
	userValidate   validates.UserValidate
	authMiddleware middlewares.AuthMiddleware
	userMiddleware middlewares.UserMiddleware
}

func NewUserRoutes(
	app fiber.Router,
	userHandler handlers.UserHandler,
	userValidate validates.UserValidate,
	authMiddleware middlewares.AuthMiddleware,
	userMiddleware middlewares.UserMiddleware,
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
	users.Get("/", r.userValidate.ValidateGetUsersRequest, r.userHandler.GetUsers)

	// Method POST
	users.Post("/", r.userValidate.ValidateCreateUserRequest, r.userHandler.CreateUser)

	// Method PATCH
	users.Patch("/:id", r.userMiddleware.UpdateUser, r.userValidate.ValidateUpdateUserRequest, r.userHandler.UpdateUser)
}
