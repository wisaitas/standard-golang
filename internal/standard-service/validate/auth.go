package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/share-pkg/validator"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
)

type AuthValidate interface {
	LoginRequest(c *fiber.Ctx) error
	RegisterRequest(c *fiber.Ctx) error
}

type authValidate struct {
	validator validator.Validator
}

func NewAuthValidate(
	validator validator.Validator,
) AuthValidate {
	return &authValidate{
		validator: validator,
	}
}

func (v *authValidate) LoginRequest(c *fiber.Ctx) error {
	req := request.LoginRequest{}

	if err := v.validator.ValidateCommonJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (v *authValidate) RegisterRequest(c *fiber.Ctx) error {
	req := request.RegisterRequest{}

	if err := v.validator.ValidateCommonJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	c.Locals("req", req)
	return c.Next()
}
