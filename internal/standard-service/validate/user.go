package validate

import (
	"github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/validator"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/param"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"

	"github.com/gofiber/fiber/v2"
)

type UserValidate interface {
	CreateUser(c *fiber.Ctx) error
	GetUsers(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
}

type userValidate struct {
	validator validator.Validator
}

func NewUserValidate(
	validator validator.Validator,
) UserValidate {
	return &userValidate{
		validator: validator,
	}
}

func (v *userValidate) CreateUser(c *fiber.Ctx) error {
	req := request.CreateUserRequest{}

	if err := v.validator.ValidateCommonJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (v *userValidate) GetUsers(c *fiber.Ctx) error {
	query := repository.PaginationQuery{}

	if err := v.validator.ValidateCommonQueryParam(c, &query); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("query", query)
	return c.Next()

}

func (v *userValidate) UpdateUser(c *fiber.Ctx) error {
	req := request.UpdateUserRequest{}
	params := param.UserParam{}

	if err := v.validator.ValidateCommonParam(c, &params); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	if err := v.validator.ValidateCommonJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("req", req)
	c.Locals("params", params)
	return c.Next()
}
