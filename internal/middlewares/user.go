package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type UserMiddleware struct {
	redisUtil pkg.RedisClient
}

func NewUserMiddleware(redisUtil pkg.RedisClient) *UserMiddleware {
	return &UserMiddleware{
		redisUtil: redisUtil,
	}
}

func (r *UserMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	// if handler permission
	// userContext, ok := c.Locals("userContext").(models.UserContext)
	// if !ok {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
	// 		Message: "user context not found",
	// 	})
	// }

	// if userContext.Username != "test" {
	// 	return c.Status(fiber.StatusForbidden).JSON(pkg.ErrorResponse{
	// 		Message: "you are not authorized to access this resource",
	// 	})
	// }

	return c.Next()
}
