package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handlers"
	"github.com/wisaitas/standard-golang/internal/standard-service/validates"
)

type ProvinceRoutes struct {
	app              fiber.Router
	provinceHandler  handlers.ProvinceHandler
	provinceValidate validates.ProvinceValidate
}

func NewProvinceRoutes(
	app fiber.Router,
	provinceHandler handlers.ProvinceHandler,
	provinceValidate validates.ProvinceValidate,
) *ProvinceRoutes {
	return &ProvinceRoutes{
		app:              app,
		provinceHandler:  provinceHandler,
		provinceValidate: provinceValidate,
	}
}

func (r *ProvinceRoutes) ProvinceRoutes() {
	provinces := r.app.Group("/provinces")

	// Method GET
	provinces.Get("/", r.provinceValidate.ValidateGetProvincesRequest, r.provinceHandler.GetProvinces)
}
