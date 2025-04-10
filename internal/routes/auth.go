package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/handlers"
	"github.com/wisaitas/standard-golang/internal/middlewares"
	"github.com/wisaitas/standard-golang/internal/validates"
)

type AuthRoutes struct {
	app            fiber.Router
	authHandler    *handlers.AuthHandler
	authValidate   validates.AuthValidate
	authMiddleware middlewares.AuthMiddleware
}

func NewAuthRoutes(
	app fiber.Router,
	authHandler *handlers.AuthHandler,
	authValidate validates.AuthValidate,
	authMiddleware middlewares.AuthMiddleware,

) *AuthRoutes {
	return &AuthRoutes{
		app:            app,
		authHandler:    authHandler,
		authValidate:   authValidate,
		authMiddleware: authMiddleware,
	}
}

func (r *AuthRoutes) AuthRoutes() {
	auth := r.app.Group("/auth")

	// Method POST
	auth.Post("/login", r.authValidate.ValidateLoginRequest, r.authHandler.Login)
	auth.Post("/logout", r.authMiddleware.Logout, r.authHandler.Logout)
	auth.Post("/register", r.authValidate.ValidateRegisterRequest, r.authHandler.Register)
	auth.Post("/refresh-token", r.authMiddleware.RefreshToken, r.authHandler.RefreshToken)
}
