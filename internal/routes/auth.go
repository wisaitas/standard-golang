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
	authValidate   *validates.AuthValidate
	authMiddleware *middlewares.AuthMiddleware
}

func NewAuthRoutes(
	app fiber.Router,
	authHandler *handlers.AuthHandler,
	authValidate *validates.AuthValidate,
	authMiddleware *middlewares.AuthMiddleware,

) *AuthRoutes {
	return &AuthRoutes{
		app:            app,
		authHandler:    authHandler,
		authValidate:   authValidate,
		authMiddleware: authMiddleware,
	}
}

func (r *AuthRoutes) AuthRoutes() {
	// Method POST
	r.app.Post("/login", r.authValidate.ValidateLoginRequest, r.authHandler.Login)
	r.app.Post("/logout", r.authMiddleware.Logout, r.authHandler.Logout)
	r.app.Post("/register", r.authValidate.ValidateRegisterRequest, r.authHandler.Register)
	r.app.Post("/refresh-token", r.authMiddleware.RefreshToken, r.authHandler.RefreshToken)
}
