package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthValidate interface {
	ValidateLoginRequest(c *fiber.Ctx) error
	ValidateRegisterRequest(c *fiber.Ctx) error
}

type authValidate struct {
	validator pkg.ValidatorUtil
}

func NewAuthValidate(
	validator pkg.ValidatorUtil,
) AuthValidate {
	return &authValidate{
		validator: validator,
	}
}

func (r *authValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := requests.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *authValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := requests.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
