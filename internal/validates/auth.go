package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthValidate struct {
	validator pkg.ValidatorUtil
}

func NewAuthValidate(
	validator pkg.ValidatorUtil,
) *AuthValidate {
	return &AuthValidate{
		validator: validator,
	}
}

func (r *AuthValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := requests.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *AuthValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := requests.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req, r.validator); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
