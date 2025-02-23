package validates

import (
	"fmt"

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

	if err := validateCommonRequestJSONBody(c, &req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("req", req)
	return c.Next()
}

func (r *UserValidate) ValidateGetUsersRequest(c *fiber.Ctx) error {
	querys := request.PaginationQuery{}

	if err := validateCommonPaginationQuery(c, &querys); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: fmt.Sprintf("failed to validate request: %s", err.Error()),
		})
	}

	c.Locals("querys", querys)
	return c.Next()

}
