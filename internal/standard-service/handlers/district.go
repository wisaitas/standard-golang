package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/queries"
	districtService "github.com/wisaitas/standard-golang/internal/standard-service/services/district"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictHandler interface {
	GetDistricts(c *fiber.Ctx) error
}

type districtHandler struct {
	districtService districtService.DistrictService
}

func NewDistrictHandler(
	districtService districtService.DistrictService,
) DistrictHandler {
	return &districtHandler{
		districtService: districtService,
	}
}

// @Summary Get Districts
// @Description Get Districts
// @Tags District
// @Accept json
// @Produce json
// @Param query query queries.DistrictQuery true "District Query"
// @Success 200 {object} pkg.SuccessResponse
// @Router /districts [get]
func (r *districtHandler) GetDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.DistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	districts, statusCode, err := r.districtService.GetDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "Districts fetched successfully",
		Data:    districts,
	})
}
