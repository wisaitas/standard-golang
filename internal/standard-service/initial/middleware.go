package initial

import (
	"github.com/gofiber/fiber/v2"
	middlewareInternal "github.com/wisaitas/standard-golang/internal/standard-service/middleware"
	middlewareConfig "github.com/wisaitas/standard-golang/internal/standard-service/middleware/configs"
)

type middleware struct {
	AuthMiddleware middlewareInternal.AuthMiddleware
	UserMiddleware middlewareInternal.UserMiddleware
}

func newMiddleware(lib *lib) *middleware {
	return &middleware{
		AuthMiddleware: middlewareInternal.NewAuthMiddleware(lib.redis, lib.jwt),
		UserMiddleware: middlewareInternal.NewUserMiddleware(lib.redis, lib.jwt),
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
