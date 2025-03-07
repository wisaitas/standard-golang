package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	subDistrictService "github.com/wisaitas/standard-golang/internal/services/sub-district"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type SubDistrictHandler struct {
	subDistrictService subDistrictService.SubDistrictService
}

func NewSubDistrictHandler(
	subDistrictService subDistrictService.SubDistrictService,
) *SubDistrictHandler {
	return &SubDistrictHandler{
		subDistrictService: subDistrictService,
	}
}

func (r *SubDistrictHandler) GetSubDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(queries.SubDistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(errors.New("failed to get queries")).Error(),
		})
	}

	subDistricts, statusCode, err := r.subDistrictService.GetSubDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Status(statusCode).JSON(responses.SuccessResponse{
		Message: "SubDistricts fetched successfully",
		Data:    subDistricts,
	})
}
