package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictValidate struct {
	validator pkg.ValidatorUtil
}

func NewDistrictValidate(
	validator pkg.ValidatorUtil,
) *DistrictValidate {
	return &DistrictValidate{
		validator: validator,
	}
}

func (r *DistrictValidate) ValidateGetDistrictsRequest(c *fiber.Ctx) error {
	query := queries.DistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
