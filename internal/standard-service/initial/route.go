package initial

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/internal/standard-service/routes"
)

type route struct {
	UserRoutes        *routes.UserRoutes
	AuthRoutes        *routes.AuthRoutes
	ProvinceRoutes    *routes.ProvinceRoutes
	DistrictRoutes    *routes.DistrictRoutes
	SubDistrictRoutes *routes.SubDistrictRoutes
}

func newRoute(
	app *fiber.App,
	handler *handler,
	validate *validate,
	middleware *middleware,
) {
	if env.ENV == constants.Util.Dev {
		app.Get("/swagger/*", swagger.New(
			swagger.Config{},
		))
	}
	apiRoute := app.Group("/api/v1")

	route := route{
		UserRoutes: routes.NewUserRoutes(
			apiRoute,
			handler.userHandler,
			validate.userValidate,
			middleware.AuthMiddleware,
			middleware.UserMiddleware,
		),
		AuthRoutes: routes.NewAuthRoutes(
			apiRoute,
			handler.authHandler,
			validate.authValidate,
			middleware.AuthMiddleware,
		),
		ProvinceRoutes: routes.NewProvinceRoutes(
			apiRoute,
			handler.provinceHandler,
			validate.provinceValidate,
		),
		DistrictRoutes: routes.NewDistrictRoutes(
			apiRoute,
			handler.districtHandler,
			validate.districtValidate,
		),
		SubDistrictRoutes: routes.NewSubDistrictRoutes(
			apiRoute,
			handler.subDistrictHandler,
			validate.subDistrictValidate,
		),
	}

	route.setupRoute()
}

func (r *route) setupRoute() {
	r.UserRoutes.UserRoutes()
	r.AuthRoutes.AuthRoutes()
	r.ProvinceRoutes.ProvinceRoutes()
	r.DistrictRoutes.DistrictRoutes()
	r.SubDistrictRoutes.SubDistrictRoutes()
}
