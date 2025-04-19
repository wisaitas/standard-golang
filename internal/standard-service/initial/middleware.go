package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/middlewares"
	middlewareConfig "github.com/wisaitas/standard-golang/internal/standard-service/middlewares/configs"
)

type middleware struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

func newMiddleware(util *util) *middleware {
	return &middleware{
		AuthMiddleware: middlewares.NewAuthMiddleware(util.redisUtil, util.jwtUtil),
		UserMiddleware: middlewares.NewUserMiddleware(util.redisUtil, util.jwtUtil),
	}
}

func setupMiddleware(app *fiber.App) {
	app.Use(
		middlewareConfig.Recovery(),
		middlewareConfig.Limiter(),
		middlewareConfig.CORS(),
		middlewareConfig.Logger(),
		middlewareConfig.Healthz(),
	)
}
