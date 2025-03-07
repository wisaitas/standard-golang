package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
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
			Message: utils.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()
}
