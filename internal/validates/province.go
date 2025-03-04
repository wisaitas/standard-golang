package validates

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
)

type ProvinceValidate struct {
}

func NewProvinceValidate() *ProvinceValidate {
	return &ProvinceValidate{}
}

func (r *ProvinceValidate) ValidateGetProvincesRequest(c *fiber.Ctx) error {
	query := queries.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
