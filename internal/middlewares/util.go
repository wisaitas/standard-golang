package middlewares

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/configs"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/utils"
)

func authToken(c *fiber.Ctx, redisUtil utils.RedisClient) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return errors.New("invalid token type")
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
		return err
	}

	_, err = redisUtil.Get(context.Background(), fmt.Sprintf("access_token:%s", uuid.MustParse(userContext.ID)))
	if err != nil {
		if err == redis.Nil {
			return errors.New("session not found")
		}

		return err
	}

	c.Locals("userContext", userContext)
	return nil

}
