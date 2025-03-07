package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	districtService "github.com/wisaitas/standard-golang/internal/services/district"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type DistrictHandler struct {
	districtService districtService.DistrictService
}

func NewDistrictHandler(
	districtService districtService.DistrictService,
) *DistrictHandler {
	return &DistrictHandler{
		districtService: districtService,
	}
}

func (r *DistrictHandler) GetDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.DistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(errors.New("failed to get queries")).Error(),
		})
	}

	districts, statusCode, err := r.districtService.GetDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "Districts fetched successfully",
		Data:    districts,
	})
}
