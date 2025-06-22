package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/utils"
	provinceService "github.com/wisaitas/standard-golang/internal/standard-service/service/province"
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
	query, ok := c.Locals("query").(repositoryPkg.PaginationQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get queries")),
		})
	}

	provinces, statusCode, err := r.provinceService.GetProvinces(query)
	if err != nil {
		return c.Status(statusCode).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(response.ApiResponse[any]{
		Message: "Provinces fetched successfully",
		Data:    provinces,
	})
}
