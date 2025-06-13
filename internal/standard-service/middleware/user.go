package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/pkg"
)

type UserMiddleware interface {
	UpdateUser(c *fiber.Ctx) error
}

type userMiddleware struct {
	redis pkg.Redis
	jwt   pkg.JWT
}

func NewUserMiddleware(redis pkg.Redis, jwt pkg.JWT) UserMiddleware {
	return &userMiddleware{
		redis: redis,
		jwt:   jwt,
	}
}

func (r *userMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := r.jwt.AuthAccessToken(c, r.redis, r.jwt, env.JWT_SECRET); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: err.Error(),
		})
	}

	return c.Next()
}
