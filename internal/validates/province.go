package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type ProvinceValidate struct {
	validator pkg.ValidatorUtil
}

func NewProvinceValidate(
	validator pkg.ValidatorUtil,
) *ProvinceValidate {
	return &ProvinceValidate{
		validator: validator,
	}
}

func (r *ProvinceValidate) ValidateGetProvincesRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
