package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handler"
	"github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type ProvinceRoutes struct {
	app              fiber.Router
	provinceHandler  handler.ProvinceHandler
	provinceValidate validate.ProvinceValidate
}

func NewProvinceRoutes(
	app fiber.Router,
	provinceHandler handler.ProvinceHandler,
	provinceValidate validate.ProvinceValidate,
) *ProvinceRoutes {
	return &ProvinceRoutes{
		app:              app,
		provinceHandler:  provinceHandler,
		provinceValidate: provinceValidate,
	}
}

func (r *ProvinceRoutes) ProvinceRoutes() {
	provinces := r.app.Group("/provinces")

	provinces.Get("/", r.provinceValidate.GetProvinces, r.provinceHandler.GetProvinces)
}
