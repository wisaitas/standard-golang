package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/share-pkg/auth/jwt"
	"github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/share-pkg/response"
	standardservice "github.com/wisaitas/standard-golang/internal/standard-service"
)

type UserMiddleware interface {
	UpdateUser(c *fiber.Ctx) error
}

type userMiddleware struct {
	redis redis.Redis
	jwt   jwt.Jwt
}

func NewUserMiddleware(redis redis.Redis, jwt jwt.Jwt) UserMiddleware {
	return &userMiddleware{
		redis: redis,
		jwt:   jwt,
	}
}

func (r *userMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := r.jwt.AuthAccessToken(c, r.redis, r.jwt, standardservice.ENV.Server.JwtSecret); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ApiResponse[any]{
			Error: err,
		})
	}

	return c.Next()
}
