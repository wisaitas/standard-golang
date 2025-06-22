package handler

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	responsePkg "github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	subDistrictService "github.com/wisaitas/standard-golang/internal/standard-service/service/sub-district"
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

func (r *subDistrictHandler) GetSubDistricts(c *fiber.Ctx) error {
	query, ok := c.Locals("query").(query.SubDistrictQuery)
	if !ok {
		return c.Status(fiber.StatusBadRequest).JSON(responsePkg.ApiResponse[any]{
			Error: utils.Error(errors.New("failed to get queries")),
		})
	}

	subDistricts, statusCode, err := r.subDistrictService.GetSubDistricts(query)
	if err != nil {
		return c.Status(statusCode).JSON(responsePkg.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Status(statusCode).JSON(responsePkg.ApiResponse[[]response.SubDistrictResponse]{
		Message: "SubDistricts fetched successfully",
		Data:    subDistricts,
	})
}
