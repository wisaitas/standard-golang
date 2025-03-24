package pkg

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims interface {
	jwt.Claims
	GetID() string
}

type StandardClaims struct {
	jwt.RegisteredClaims
	ID string `json:"id"`
}

func (s StandardClaims) GetID() string {
	return s.ID
}

type JWTUtil interface {
	Generate(claims JWTClaims, secret string) (string, error)
	Parse(tokenString string, claims jwt.Claims, secret string) (jwt.Claims, error)
	ExtractTokenFromHeader(c *fiber.Ctx) (string, error)
	ValidateToken(tokenString string, claims jwt.Claims, secret string) error
	CreateStandardClaims(id string, expireTime time.Duration) StandardClaims
}

type jwtUtil struct{}

func NewJWTUtil() JWTUtil {
	return &jwtUtil{}
}

func (j *jwtUtil) Generate(claims JWTClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", Error(err)
	}

	return tokenString, nil
}

func (j *jwtUtil) ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	authHeader := c.Get("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return "", Error(errors.New("invalid token type"))
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	return token, nil
}

func (j *jwtUtil) Parse(tokenString string, claims jwt.Claims, secret string) (jwt.Claims, error) {
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, Error(err)
	}

	return claims, nil
}

func (j *jwtUtil) ValidateToken(tokenString string, claims jwt.Claims, secret string) error {
	_, err := j.Parse(tokenString, claims, secret)
	return err
}

func (j *jwtUtil) CreateStandardClaims(id string, expireTime time.Duration) StandardClaims {
	return StandardClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
}
