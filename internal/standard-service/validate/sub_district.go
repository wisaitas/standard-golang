package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/validator"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
)

type SubDistrictValidate interface {
	GetSubDistricts(c *fiber.Ctx) error
}

type subDistrictValidate struct {
	validator validator.Validator
}

func NewSubDistrictValidate(
	validator validator.Validator,
) SubDistrictValidate {
	return &subDistrictValidate{
		validator: validator,
	}
}

func (v *subDistrictValidate) GetSubDistricts(c *fiber.Ctx) error {
	query := query.SubDistrictQuery{}

	if err := v.validator.ValidateCommonQueryParam(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("query", query)
	return c.Next()
}
