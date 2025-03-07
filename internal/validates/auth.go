package validates

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type AuthValidate struct {
}

func NewAuthValidate() *AuthValidate {
	return &AuthValidate{}
}

func (r *AuthValidate) ValidateLoginRequest(c *fiber.Ctx) error {
	req := requests.LoginRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *AuthValidate) ValidateRegisterRequest(c *fiber.Ctx) error {
	req := requests.RegisterRequest{}

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
