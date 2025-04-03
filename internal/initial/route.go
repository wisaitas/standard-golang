package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/routes"
)

type Routes struct {
	UserRoutes        *routes.UserRoutes
	AuthRoutes        *routes.AuthRoutes
	ProvinceRoutes    *routes.ProvinceRoutes
	DistrictRoutes    *routes.DistrictRoutes
	SubDistrictRoutes *routes.SubDistrictRoutes
}

func initializeRoutes(
	apiRoutes fiber.Router,
	handlers *Handler,
	validates *Validate,
	middlewares *Middleware,
) *Routes {
	return &Routes{
		UserRoutes: routes.NewUserRoutes(
			apiRoutes,
			&handlers.UserHandler,
			validates.UserValidate,
			middlewares.AuthMiddleware,
			middlewares.UserMiddleware,
		),
		AuthRoutes: routes.NewAuthRoutes(
			apiRoutes,
			&handlers.AuthHandler,
			validates.AuthValidate,
			middlewares.AuthMiddleware,
		),
		ProvinceRoutes: routes.NewProvinceRoutes(
			apiRoutes,
			&handlers.ProvinceHandler,
			&validates.ProvinceValidate,
		),
		DistrictRoutes: routes.NewDistrictRoutes(
			apiRoutes,
			&handlers.DistrictHandler,
			&validates.DistrictValidate,
		),
		SubDistrictRoutes: routes.NewSubDistrictRoutes(
			apiRoutes,
			&handlers.SubDistrictHandler,
			&validates.SubDistrictValidate,
		),
	}
}

func (r *Routes) SetupRoutes() {
	r.UserRoutes.UserRoutes()
	r.AuthRoutes.AuthRoutes()
	r.ProvinceRoutes.ProvinceRoutes()
	r.DistrictRoutes.DistrictRoutes()
	r.SubDistrictRoutes.SubDistrictRoutes()
}
