package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handler"
	"github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type DistrictRoutes struct {
	app              fiber.Router
	districtHandler  handler.DistrictHandler
	districtValidate validate.DistrictValidate
}

func NewDistrictRoutes(
	app fiber.Router,
	districtHandler handler.DistrictHandler,
	districtValidate validate.DistrictValidate,
) *DistrictRoutes {
	return &DistrictRoutes{
		app:              app,
		districtHandler:  districtHandler,
		districtValidate: districtValidate,
	}
}

func (r *DistrictRoutes) DistrictRoutes() {
	districts := r.app.Group("/districts")

	districts.Get("/", r.districtValidate.GetDistricts, r.districtHandler.GetDistricts)
}
