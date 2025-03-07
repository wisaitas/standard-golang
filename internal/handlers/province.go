package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	provinceService "github.com/wisaitas/standard-golang/internal/services/province"
	"github.com/wisaitas/standard-golang/internal/utils"
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
			Message: utils.Error(errors.New("failed to get queries")).Error(),
		})
	}

	provinces, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responses.SuccessResponse{
		Message: "Provinces fetched successfully",
		Data:    provinces,
	})
}
