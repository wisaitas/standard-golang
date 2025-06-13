package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictValidate interface {
	GetSubDistricts(c *fiber.Ctx) error
}

type subDistrictValidate struct {
	validator pkg.Validator
}

func NewSubDistrictValidate(
	validator pkg.Validator,
) SubDistrictValidate {
	return &subDistrictValidate{
		validator: validator,
	}
}

func (v *subDistrictValidate) GetSubDistricts(c *fiber.Ctx) error {
	query := query.SubDistrictQuery{}

	if err := v.validator.ValidateCommonQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
