package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type TokenContext struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.RegisteredClaims
}

type UserContext struct {
	UserID       uuid.UUID `json:"user_id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	BirthDate    time.Time `json:"birth_date"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
