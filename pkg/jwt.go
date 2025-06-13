package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	jwtLib "github.com/golang-jwt/jwt/v5"
	redisLib "github.com/redis/go-redis/v9"
)

type JWTClaims interface {
	jwtLib.Claims
	GetID() string
}

type StandardClaims struct {
	jwtLib.RegisteredClaims
	ID string `json:"id"`
}

func (s StandardClaims) GetID() string {
	return s.ID
}

type JWT interface {
	Generate(claims JWTClaims, secret string) (string, error)
	Parse(tokenString string, claims jwtLib.Claims, secret string) (jwtLib.Claims, error)
	ExtractTokenFromHeader(c *fiber.Ctx) (string, error)
	ValidateToken(tokenString string, claims jwtLib.Claims, secret string) error
	CreateStandardClaims(id string, expireTime time.Duration) StandardClaims
	AuthAccessToken(c *fiber.Ctx, redis Redis, jwt JWT, secret string) error
	AuthRefreshToken(c *fiber.Ctx, redis Redis, jwt JWT, secret string) error
}

type jwt struct{}

func NewJWT() JWT {
	return &jwt{}
}

func (j *jwt) Generate(claims JWTClaims, secret string) (string, error) {
	token := jwtLib.NewWithClaims(jwtLib.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", fmt.Errorf("[jwt] %w", err)
	}

	return tokenString, nil
}

func (j *jwt) ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("[jwt] invalid token type")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}

func (j *jwt) Parse(tokenString string, claims jwtLib.Claims, secret string) (jwtLib.Claims, error) {
	_, err := jwtLib.ParseWithClaims(tokenString, claims, func(token *jwtLib.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtLib.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[jwt] unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("[jwt] %w", err)
	}

	return claims, nil
}

func (j *jwt) ValidateToken(tokenString string, claims jwtLib.Claims, secret string) error {
	_, err := j.Parse(tokenString, claims, secret)
	return err
}

func (j *jwt) CreateStandardClaims(id string, expireTime time.Duration) StandardClaims {
	return StandardClaims{
		ID: id,
		RegisteredClaims: jwtLib.RegisteredClaims{
			ExpiresAt: jwtLib.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwtLib.NewNumericDate(time.Now()),
		},
	}
}

func (j *jwt) AuthAccessToken(c *fiber.Ctx, redis Redis, jwt JWT, secret string) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fmt.Errorf("[jwt] invalid token type")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var tokenContext TokenContext
	_, err := jwt.Parse(token, &tokenContext, secret)
	if err != nil {
		return fmt.Errorf("[jwt] %w", err)
	}

	userContextJSON, err := redis.Get(context.Background(), fmt.Sprintf("access_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redisLib.Nil {
			return fmt.Errorf("[jwt] session not found")
		}

		return fmt.Errorf("[jwt] %w", err)
	}

	var userContext UserContext
	if err := json.Unmarshal([]byte(userContextJSON), &userContext); err != nil {
		return fmt.Errorf("[jwt] %w", err)
	}

	c.Locals("userContext", userContext)
	return nil
}

func (j *jwt) AuthRefreshToken(c *fiber.Ctx, redis Redis, jwt JWT, secret string) error {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return fmt.Errorf("[jwt] invalid token type")
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	var tokenContext TokenContext
	_, err := jwt.Parse(token, &tokenContext, secret)
	if err != nil {
		return fmt.Errorf("[jwt] %w", err)
	}

	userContextJSON, err := redis.Get(context.Background(), fmt.Sprintf("refresh_token:%s", tokenContext.UserID))
	if err != nil {
		if err == redisLib.Nil {
			return fmt.Errorf("[jwt] session not found")
		}

		return fmt.Errorf("[jwt] %w", err)
	}

	var userContext UserContext
	if err := json.Unmarshal([]byte(userContextJSON), &userContext); err != nil {
		return fmt.Errorf("[jwt] %w", err)
	}

	c.Locals("userContext", userContext)
	return nil
}
