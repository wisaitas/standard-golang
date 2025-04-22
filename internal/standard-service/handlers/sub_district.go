package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/queries"
	subDistrictService "github.com/wisaitas/standard-golang/internal/standard-service/services/sub-district"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictHandler interface {
	GetSubDistricts(c *fiber.Ctx) error
}

type subDistrictHandler struct {
	subDistrictService subDistrictService.SubDistrictService
}

func NewSubDistrictHandler(
	subDistrictService subDistrictService.SubDistrictService,
) SubDistrictHandler {
	return &subDistrictHandler{
		subDistrictService: subDistrictService,
	}
}

// @Summary Get Sub Districts
// @Description Get Sub Districts
// @Tags Sub District
// @Accept json
// @Produce json
// @Param query query queries.SubDistrictQuery true "Sub District Query"
// @Success 200 {object} pkg.SuccessResponse
// @Router /sub-districts [get]
func (r *subDistrictHandler) GetSubDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.SubDistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	subDistricts, statusCode, err := r.subDistrictService.GetSubDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(pkg.SuccessResponse{
		Message: "SubDistricts fetched successfully",
		Data:    subDistricts,
	})
}
