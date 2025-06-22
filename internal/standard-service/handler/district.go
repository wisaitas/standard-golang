package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
	districtService "github.com/wisaitas/standard-golang/internal/standard-service/service/district"
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

func (r *districtHandler) GetDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(query.DistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get queries")),
		})
	}

	districts, statusCode, err := r.districtService.GetDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(response.ApiResponse[any]{
		Message: "Districts fetched successfully",
		Data:    districts,
	})
}
