package pkg

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/mock"
	"github.com/wisaitas/standard-golang/pkg"
)

type MockJwt struct {
	mock.Mock
}

func NewMockJwt() *MockJwt {
	return &MockJwt{}
}

func (r *MockJwt) Generate(claims pkg.JWTClaims, secret string) (string, error) {
	args := r.Called(claims, secret)
	return args.String(0), args.Error(1)
}

func (r *MockJwt) Parse(tokenString string, claims jwt.Claims, secret string) (jwt.Claims, error) {
	args := r.Called(tokenString, claims, secret)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return claims, args.Error(1)
}

func (r *MockJwt) ExtractTokenFromHeader(c *fiber.Ctx) (string, error) {
	args := r.Called(c)
	return args.String(0), args.Error(1)
}

func (r *MockJwt) ValidateToken(tokenString string, claims jwt.Claims, secret string) error {
	args := r.Called(tokenString, claims, secret)
	return args.Error(0)
}

func (r *MockJwt) CreateStandardClaims(id string, expireTime time.Duration) pkg.StandardClaims {
	args := r.Called(id, expireTime)
	return args.Get(0).(pkg.StandardClaims)
}

func (r *MockJwt) AuthAccessToken(c *fiber.Ctx, redis pkg.Redis, jwt pkg.JWT, secret string) error {
	args := r.Called(c, redis, jwt, secret)
	return args.Error(0)
}

func (r *MockJwt) AuthRefreshToken(c *fiber.Ctx, redis pkg.Redis, jwt pkg.JWT, secret string) error {
	args := r.Called(c, redis, jwt, secret)
	return args.Error(0)
}

func (r *MockJwt) GenerateToken(data map[string]interface{}, exp int64, secret string) (string, error) {
	args := r.Called(data, exp, secret)
	return args.String(0), args.Error(1)
}

func (r *MockJwt) ExtractClaims(tokenString string) (map[string]interface{}, error) {
	args := r.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(map[string]interface{}), args.Error(1)
}

func (r *MockJwt) ParseToken(tokenString string) (*jwt.Token, error) {
	args := r.Called(tokenString)

	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*jwt.Token), args.Error(1)
}
