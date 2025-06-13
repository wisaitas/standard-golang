package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictValidate interface {
	GetDistricts(c *fiber.Ctx) error
}

type districtValidate struct {
	validator pkg.Validator
}

func NewDistrictValidate(
	validator pkg.Validator,
) DistrictValidate {
	return &districtValidate{
		validator: validator,
	}
}

func (v *districtValidate) GetDistricts(c *fiber.Ctx) error {
	query := query.DistrictQuery{}

	if err := v.validator.ValidateCommonQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
