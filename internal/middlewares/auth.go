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
	"github.com/wisaitas/standard-golang/internal/dtos/response"
	"github.com/wisaitas/standard-golang/internal/models"
)

type AuthMiddleware struct {
	redis *redis.Client
}

func NewAuthMiddleware(redis *redis.Client) *AuthMiddleware {
	return &AuthMiddleware{
		redis: redis,
	}
}

func (r *AuthMiddleware) AuthToken(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
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
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	_, err = r.redis.Get(context.Background(), fmt.Sprintf("access_token:%s", uuid.MustParse(userContext.ID))).Result()
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: err.Error(),
		})
	}

	c.Locals("userContext", userContext)
	return c.Next()

}
