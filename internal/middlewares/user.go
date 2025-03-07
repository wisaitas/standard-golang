package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type UserMiddleware struct {
	redisUtil utils.RedisClient
}

func NewUserMiddleware(redisUtil utils.RedisClient) *UserMiddleware {
	return &UserMiddleware{
		redisUtil: redisUtil,
	}
}

func (r *UserMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	// if handler permission
	// userContext, ok := c.Locals("userContext").(models.UserContext)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
	// 		Message: "user context not found",
	// 	})
	// }

	// if userContext.Username != "test" {
	// 	return c.Status(fiber.StatusForbidden).JSON(responses.ErrorResponse{
	// 		Message: "you are not authorized to access this resource",
	// 	})
	// }

	return c.Next()
}
