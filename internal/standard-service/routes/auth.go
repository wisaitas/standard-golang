package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handler"
	"github.com/wisaitas/standard-golang/internal/standard-service/middleware"
	"github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type AuthRoutes struct {
	app            fiber.Router
	authHandler    handler.AuthHandler
	authValidate   validate.AuthValidate
	authMiddleware middleware.AuthMiddleware
}

func NewAuthRoutes(
	app fiber.Router,
	authHandler handler.AuthHandler,
	authValidate validate.AuthValidate,
	authMiddleware middleware.AuthMiddleware,

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
	auth.Post("/login", r.authValidate.LoginRequest, r.authHandler.Login)
	auth.Post("/logout", r.authMiddleware.Logout, r.authHandler.Logout)
	auth.Post("/register", r.authValidate.RegisterRequest, r.authHandler.Register)
	auth.Post("/refresh-token", r.authMiddleware.RefreshToken, r.authHandler.RefreshToken)
}
