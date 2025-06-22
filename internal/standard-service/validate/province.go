package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/validator"
)

type ProvinceValidate interface {
	GetProvinces(c *fiber.Ctx) error
}

type provinceValidate struct {
	validator validator.Validator
}

func NewProvinceValidate(
	validator validator.Validator,
) ProvinceValidate {
	return &provinceValidate{
		validator: validator,
	}
}

func (v *provinceValidate) GetProvinces(c *fiber.Ctx) error {
	query := repository.PaginationQuery{}

	if err := v.validator.ValidateCommonQueryParam(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("query", query)
	return c.Next()
}
