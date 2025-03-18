package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthMiddleware struct {
	redisUtil pkg.RedisClient
}

func NewAuthMiddleware(redisUtil pkg.RedisClient) *AuthMiddleware {
	return &AuthMiddleware{
		redisUtil: redisUtil,
	}
}

func (r *AuthMiddleware) Logout(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()

}

func (r *AuthMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := authToken(c, r.redisUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}
