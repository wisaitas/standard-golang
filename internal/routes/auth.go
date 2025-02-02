package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/handlers"
	"github.com/wisaitas/standard-golang/internal/validates"
)

type AuthRoutes struct {
	app          fiber.Router
	authHandler  *handlers.AuthHandler
	authValidate *validates.AuthValidate
}

func NewAuthRoutes(
	app fiber.Router,
	authHandler *handlers.AuthHandler,
	authValidate *validates.AuthValidate,
) *AuthRoutes {
	return &AuthRoutes{
		app:          app,
		authHandler:  authHandler,
		authValidate: authValidate,
	}
}

func (r *AuthRoutes) AuthRoutes() {
	r.app.Post("/login", r.authValidate.ValidateLoginRequest, r.authHandler.Login)
	r.app.Post("/register", r.authValidate.ValidateRegisterRequest, r.authHandler.Register)
}
