package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/routes"
)

type Routes struct {
	UserRoutes *routes.UserRoutes
	AuthRoutes *routes.AuthRoutes
}

func InitializeRoutes(
	apiRoutes fiber.Router,
	handlers *Handlers,
	validates *Validates,
	middlewares *Middlewares,
) *Routes {
	return &Routes{
		UserRoutes: routes.NewUserRoutes(
			apiRoutes,
			&handlers.UserHandler,
			&validates.UserValidate,
			&middlewares.AuthMiddleware,
			&middlewares.UserMiddleware,
		),
		AuthRoutes: routes.NewAuthRoutes(
			apiRoutes,
			&handlers.AuthHandler,
			&validates.AuthValidate,
			&middlewares.AuthMiddleware,
		),
	}
}

func (r *Routes) SetupRoutes() {
	r.UserRoutes.UserRoutes()
	r.AuthRoutes.AuthRoutes()
}
