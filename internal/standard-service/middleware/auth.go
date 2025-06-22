package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/auth/jwt"
	"github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
)

type AuthMiddleware interface {
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}

type authMiddleware struct {
	redis redis.Redis
	jwt   jwt.Jwt
}

func NewAuthMiddleware(
	redis redis.Redis,
	jwt jwt.Jwt,
) AuthMiddleware {
	return &authMiddleware{
		redis: redis,
		jwt:   jwt,
	}
}

func (r *authMiddleware) Logout(c *fiber.Ctx) error {
	if err := r.jwt.AuthAccessToken(c, r.redis, r.jwt, env.Environment.Server.JwtSecret); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Next()

}

func (r *authMiddleware) RefreshToken(c *fiber.Ctx) error {
	if err := r.jwt.AuthRefreshToken(c, r.redis, r.jwt, env.Environment.Server.JwtSecret); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Next()
}
