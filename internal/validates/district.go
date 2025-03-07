package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type DistrictValidate struct {
}

func NewDistrictValidate() *DistrictValidate {
	return &DistrictValidate{}
}

func (r *DistrictValidate) ValidateGetDistrictsRequest(c *fiber.Ctx) error {
	query := queries.DistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
