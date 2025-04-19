package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wisaitas/standard-golang/pkg"
)

type UserMiddleware interface {
	UpdateUser(c *fiber.Ctx) error
}

type userMiddleware struct {
	redisUtil pkg.RedisUtil
	jwtUtil   pkg.JWTUtil
}

func NewUserMiddleware(redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) UserMiddleware {
	return &userMiddleware{
		redisUtil: redisUtil,
		jwtUtil:   jwtUtil,
	}
}

func (r *userMiddleware) UpdateUser(c *fiber.Ctx) error {
	if err := authAccessToken(c, r.redisUtil, r.jwtUtil); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(pkg.ErrorResponse{
			Message: pkg.Error(err).Error(),
		})
	}

	return c.Next()
}
