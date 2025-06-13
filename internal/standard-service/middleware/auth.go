package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/pkg"
)

type AuthMiddleware interface {
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authMiddleware struct {
	redis pkg.Redis
	jwt   pkg.JWT
}

func NewAuthMiddleware(
	redis pkg.Redis,
	jwt pkg.JWT,
) AuthMiddleware {
	return &authMiddleware{
		redis: redis,
		jwt:   jwt,
	}
}

func (r *authMiddleware) Logout(c *fiber.Ctx) error {
	if err := r.jwt.AuthAccessToken(c, r.redis, r.jwt, env.JWT_SECRET); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Next()

}

func (r *authMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := r.jwt.AuthRefreshToken(c, r.redis, r.jwt, env.JWT_SECRET); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Next()
}
