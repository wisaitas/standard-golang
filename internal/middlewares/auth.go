package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type AuthMiddleware struct {
	redisUtil utils.RedisClient
}

func NewAuthMiddleware(redisUtil utils.RedisClient) *AuthMiddleware {
	return &AuthMiddleware{
		redisUtil: redisUtil,
	}
}

func (r *AuthMiddleware) Logout(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Next()

}

func (r *AuthMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: utils.Error(err).Error(),
		})
	}

	return c.Next()
}
