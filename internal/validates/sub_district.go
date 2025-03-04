package validates

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
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
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
