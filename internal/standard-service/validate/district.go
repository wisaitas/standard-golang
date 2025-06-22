package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/validator"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
)

type DistrictValidate interface {
	GetDistricts(c *fiber.Ctx) error
}

type districtValidate struct {
	validator validator.Validator
}

func NewDistrictValidate(
	validator validator.Validator,
) DistrictValidate {
	return &districtValidate{
		validator: validator,
	}
}

func (v *districtValidate) GetDistricts(c *fiber.Ctx) error {
	query := query.DistrictQuery{}

	if err := v.validator.ValidateCommonQueryParam(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("query", query)
	return c.Next()
}
