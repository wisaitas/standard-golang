package initial

import (
	"github.com/wisaitas/standard-golang/internal/middlewares"
	"github.com/wisaitas/standard-golang/pkg"
)

type Middleware struct {
	AuthMiddleware middlewares.AuthMiddleware
	UserMiddleware middlewares.UserMiddleware
}

func NewMiddleware(redisUtil pkg.RedisUtil, jwtUtil pkg.JWTUtil) *Middleware {
	return &Middleware{
		AuthMiddleware: middlewares.NewAuthMiddleware(redisUtil, jwtUtil),
		UserMiddleware: middlewares.NewUserMiddleware(redisUtil, jwtUtil),
	}
}
