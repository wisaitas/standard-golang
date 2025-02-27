package middlewares

import (
	"context"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/configs"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type AuthMiddleware struct {
	redis utils.RedisClient
}

func NewAuthMiddleware(redis utils.RedisClient) *AuthMiddleware {
	return &AuthMiddleware{
		redis: redis,
	}
}

func (r *AuthMiddleware) AuthToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: "invalid token type",
		})
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var userContext models.UserContext
	_, err := jwt.ParseWithClaims(token, &userContext, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(configs.ENV.JWT_SECRET), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	_, err = r.redis.Get(context.Background(), fmt.Sprintf("access_token:%s", uuid.MustParse(userContext.ID)))
	if err != nil {
		if err == redis.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
				Message: "token not found",
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(responses.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("userContext", userContext)
	return c.Next()

}
