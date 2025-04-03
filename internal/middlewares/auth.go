package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthMiddleware interface {
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewAuthMiddleware(
	redisUtil pkg.RedisUtil,
	jwtUtil pkg.JWTUtil,
) AuthMiddleware {
	return &authMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *authMiddleware) Logout(c *fiber.Ctx) error {
	if err := authRefreshToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()

}

func (r *authMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := authRefreshToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}
