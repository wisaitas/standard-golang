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

func newMiddleware(sharePkg *sharePkg) *middleware {
	return &middleware{
		AuthMiddleware: middlewareInternal.NewAuthMiddleware(sharePkg.redis, sharePkg.jwt),
		UserMiddleware: middlewareInternal.NewUserMiddleware(sharePkg.redis, sharePkg.jwt),
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
