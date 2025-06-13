package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type ProvinceValidate interface {
	GetProvinces(c *fiber.Ctx) error
}

type provinceValidate struct {
	validator pkg.Validator
}

func NewProvinceValidate(
	validator pkg.Validator,
) ProvinceValidate {
	return &provinceValidate{
		validator: validator,
	}
}

func (v *provinceValidate) GetProvinces(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := v.validator.ValidateCommonQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
