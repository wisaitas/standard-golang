package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	provinceService "github.com/wisaitas/standard-golang/internal/services/province"
)

type ProvinceHandler struct {
	provinceService provinceService.ProvinceService
}

func NewProvinceHandler(
	provinceService provinceService.ProvinceService,
) *ProvinceHandler {
	return &ProvinceHandler{
		provinceService: provinceService,
	}
}

func (r *ProvinceHandler) GetProvinces(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: "failed to get queries",
		})
	}

	provinces, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
		Message: "Provinces fetched successfully",
		Data:    provinces,
	})
}
