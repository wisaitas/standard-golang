package mock_utils

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/wisaitas/standard-golang/pkg"
)

type MockJWT struct {
	mock.Mock
}

func NewMockJWT() *MockJWT {
	return &MockJWT{}
}

func (r *MockJWT) GenerateToken(userID string, role string, expiration time.Duration) (string, error) {
	args := r.Called(userID, role, expiration)
	return args.String(0), args.Error(1)
}

func (r *MockJWT) ValidateToken(tokenString string, claims jwt.Claims, secret string) error {
	args := r.Called(tokenString, claims, secret)
	return args.Error(0)
}

func (r *MockJWT) ExtractClaims(tokenString string) (map[string]interface{}, error) {
	args := r.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (r *MockJWT) ParseToken(tokenString string) (*jwt.Token, error) {
	args := r.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*jwt.Token), args.Error(1)
}

func (r *MockJWT) Parse(tokenString string, claims jwt.Claims, secret string) (jwt.Claims, error) {
	args := r.Called(tokenString, claims, secret)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return claims, args.Error(1)
}

func (r *MockJWT) CreateStandardClaims(userID string, expiration time.Duration) pkg.StandardClaims {
	args := r.Called(userID, expiration)
	return args.Get(0).(pkg.StandardClaims)
}

func (r *MockJWT) Generate(claims pkg.JWTClaims, secret string) (string, error) {
	args := r.Called(claims, secret)
	return args.String(0), args.Error(1)
}

func (r *MockJWT) ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	args := r.Called(c)
	return args.String(0), args.Error(1)
}
