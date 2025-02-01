package validates

import (
	"github.com/wisaitas/standard-golang/internal/dtos/request"
	"github.com/wisaitas/standard-golang/internal/dtos/response"

	"github.com/gofiber/fiber/v2"
)

type UserValidate struct {
}

func NewUserValidate() *UserValidate {
	return &UserValidate{}
}

func (r *UserValidate) ValidateCreateUserRequest(c *fiber.Ctx) error {
	req := request.CreateUserRequest{}

	if err := validateCommonRequest(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("req", req)
	return c.Next()
}
