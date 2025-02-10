package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/models"
)

type UserMiddleware struct {
}

func NewUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}

func (r *UserMiddleware) GetUsers(c *fiber.Ctx) error {
	userContext, ok := c.Locals("userContext").(models.UserContext)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "user context not found",
		})
	}

	if userContext.Username != "test" {
		return c.Status(fiber.StatusForbidden).JSON(response.ErrorResponse{
			Message: "you are not authorized to access this resource",
		})
	}

	return c.Next()
}
