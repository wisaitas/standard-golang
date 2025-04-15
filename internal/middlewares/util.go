package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/contexts"
	"github.com/wisaitas/standard-golang/internal/env"
	"github.com/wisaitas/standard-golang/pkg"
)

func authAccessToken(c *fiber.Ctx, redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return pkg.Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var tokenContext contexts.TokenContext
	_, err := jwtUtil.Parse(token, &tokenContext, env.JWT_SECRET)
	if err != nil {
		return pkg.Error(err)
	}

	userContextJSON, err := redisUtil.Get(context.Background(), fmt.Sprintf("access_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	var userContext contexts.UserContext
	if err := json.Unmarshal([]byte(userContextJSON), &userContext); err != nil {
		return pkg.Error(err)
	}

	c.Locals("userContext", userContext)
	return nil
}

func authRefreshToken(c *fiber.Ctx, redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return pkg.Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var tokenContext contexts.TokenContext
	_, err := jwtUtil.Parse(token, &tokenContext, env.JWT_SECRET)
	if err != nil {
		return pkg.Error(err)
	}

	userContextJSON, err := redisUtil.Get(context.Background(), fmt.Sprintf("refresh_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redis.Nil {
			return pkg.Error(errors.New("session not found"))
		}

		return pkg.Error(err)
	}

	var userContext contexts.UserContext
	if err := json.Unmarshal([]byte(userContextJSON), &userContext); err != nil {
		return pkg.Error(err)
	}

	c.Locals("userContext", userContext)
	return nil
}
