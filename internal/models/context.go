package models

import "github.com/golang-jwt/jwt/v5"

type UserContext struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
