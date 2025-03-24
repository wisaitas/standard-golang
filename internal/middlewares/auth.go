package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewAuthMiddleware(
	redisUtil pkg.RedisUtil,
	jwtUtil pkg.JWTUtil,
) *AuthMiddleware {
	return &AuthMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *AuthMiddleware) Logout(c *fiber.Ctx) error {
	if err := authRefreshToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()

}

func (r *AuthMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := authRefreshToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}
