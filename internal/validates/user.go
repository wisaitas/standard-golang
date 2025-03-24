package validates

import (
	"fmt"

	"github.com/wisaitas/standard-golang/internal/dtos/params"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/pkg"

	"github.com/gofiber/fiber/v2"
)

type UserValidate struct {
	validator pkg.ValidatorUtil
}

func NewUserValidate(
	validator pkg.ValidatorUtil,
) *UserValidate {
	return &UserValidate{
		validator: validator,
	}
}

func (r *UserValidate) ValidateCreateUserRequest(c *fiber.Ctx) error {
	req := requests.CreateUserRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *UserValidate) ValidateGetUsersRequest(c *fiber.Ctx) error {
	query := pkg.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &query, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("query", query)
	return c.Next()

}

func (r *UserValidate) ValidateUpdateUserRequest(c *fiber.Ctx) error {
	req := requests.UpdateUserRequest{}
	params := params.UserParams{}

	if err := validateCommonRequestParams(c, &params, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	c.Locals("params", params)
	return c.Next()
}
