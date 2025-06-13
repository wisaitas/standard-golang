package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handler"
	"github.com/wisaitas/standard-golang/internal/standard-service/middleware"
	"github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type UserRoutes struct {
	app            fiber.Router
	userHandler    handler.UserHandler
	userValidate   validate.UserValidate
	authMiddleware middleware.AuthMiddleware
	userMiddleware middleware.UserMiddleware
}

func NewUserRoutes(
	app fiber.Router,
	userHandler handler.UserHandler,
	userValidate validate.UserValidate,
	authMiddleware middleware.AuthMiddleware,
	userMiddleware middleware.UserMiddleware,
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
	users.Get("/", r.userValidate.GetUsers, r.userHandler.GetUsers)

	// Method POST
	users.Post("/", r.userValidate.CreateUser, r.userHandler.CreateUser)

	// Method PATCH
	users.Patch("/:id", r.userMiddleware.UpdateUser, r.userValidate.UpdateUser, r.userHandler.UpdateUser)
}
