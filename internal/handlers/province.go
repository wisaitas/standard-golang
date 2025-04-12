package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	provinceService "github.com/wisaitas/standard-golang/internal/services/province"
	"github.com/wisaitas/standard-golang/pkg"
)

type ProvinceHandler interface {
	GetProvinces(c *fiber.Ctx) error
}

type provinceHandler struct {
	provinceService provinceService.ProvinceService
}

func NewProvinceHandler(
	provinceService provinceService.ProvinceService,
) ProvinceHandler {
	return &provinceHandler{
		provinceService: provinceService,
	}
}

func (r *provinceHandler) GetProvinces(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(pkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(errors.New("failed to get queries")).Error(),
		})
	}

	provinces, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(pkg.SuccessResponse{
		Message: "Provinces fetched successfully",
		Data:    provinces,
	})
}
