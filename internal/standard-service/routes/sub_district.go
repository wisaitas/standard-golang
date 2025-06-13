package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/handler"
	"github.com/wisaitas/standard-golang/internal/standard-service/validate"
)

type SubDistrictRoutes struct {
	app                 fiber.Router
	subDistrictHandler  handler.SubDistrictHandler
	subDistrictValidate validate.SubDistrictValidate
}

func NewSubDistrictRoutes(
	app fiber.Router,
	subDistrictHandler handler.SubDistrictHandler,
	subDistrictValidate validate.SubDistrictValidate,
) *SubDistrictRoutes {
	return &SubDistrictRoutes{
		app:                 app,
		subDistrictHandler:  subDistrictHandler,
		subDistrictValidate: subDistrictValidate,
	}
}

func (r *SubDistrictRoutes) SubDistrictRoutes() {
	subDistricts := r.app.Group("/sub-districts")

	subDistricts.Get("/", r.subDistrictValidate.GetSubDistricts, r.subDistrictHandler.GetSubDistricts)
}
