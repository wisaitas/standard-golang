package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type UserMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewUserMiddleware(redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) *UserMiddleware {
	return &UserMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *UserMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := authAccessToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}
