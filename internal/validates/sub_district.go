package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictValidate struct {
	validator pkg.ValidatorUtil
}

func NewSubDistrictValidate(
	validator pkg.ValidatorUtil,
) *SubDistrictValidate {
	return &SubDistrictValidate{
		validator: validator,
	}
}

func (r *SubDistrictValidate) ValidateGetSubDistrictsRequest(c *fiber.Ctx) error {
	query := queries.SubDistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
