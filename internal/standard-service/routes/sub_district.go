package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handlers"
	"github.com/wisaitas/standard-golang/internal/standard-service/validates"
)

type SubDistrictRoutes struct {
	app                 fiber.Router
	subDistrictHandler  handlers.SubDistrictHandler
	subDistrictValidate validates.SubDistrictValidate
}

func NewSubDistrictRoutes(
	app fiber.Router,
	subDistrictHandler handlers.SubDistrictHandler,
	subDistrictValidate validates.SubDistrictValidate,
) *SubDistrictRoutes {
	return &SubDistrictRoutes{
		app:                 app,
		subDistrictHandler:  subDistrictHandler,
		subDistrictValidate: subDistrictValidate,
	}
}

func (r *SubDistrictRoutes) SubDistrictRoutes() {
	subDistricts := r.app.Group("/sub-districts")

	// Method GET
	subDistricts.Get("/", r.subDistrictValidate.ValidateGetSubDistrictsRequest, r.subDistrictHandler.GetSubDistricts)
}
