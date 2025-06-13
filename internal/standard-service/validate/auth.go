package validate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthValidate interface {
	LoginRequest(c *fiber.Ctx) error
	RegisterRequest(c *fiber.Ctx) error
}

type authValidate struct {
	validator pkg.Validator
}

func NewAuthValidate(
	validator pkg.Validator,
) AuthValidate {
	return &authValidate{
		validator: validator,
	}
}

func (v *authValidate) LoginRequest(c *fiber.Ctx) error {
	req := request.LoginRequest{}

	if err := v.validator.ValidateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (v *authValidate) RegisterRequest(c *fiber.Ctx) error {
	req := request.RegisterRequest{}

	if err := v.validator.ValidateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
