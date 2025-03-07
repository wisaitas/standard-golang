package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type SubDistrictValidate struct {
}

func NewSubDistrictValidate() *SubDistrictValidate {
	return &SubDistrictValidate{}
}

func (r *SubDistrictValidate) ValidateGetSubDistrictsRequest(c *fiber.Ctx) error {
	query := queries.SubDistrictQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
