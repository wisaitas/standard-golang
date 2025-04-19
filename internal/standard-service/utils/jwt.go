package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wisaitas/standard-golang/internal/standard-service/env"
	"github.com/wisaitas/standard-golang/pkg"
)

func GenerateToken(data map[string]interface{}, exp int64) (string, error) {
	claim := jwt.MapClaims(data)
	claim["exp"] = exp
	claim["iat"] = time.Now().Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(env.JWT_SECRET))
	if err != nil {
		return "", pkg.Error(err)
	}

	return tokenString, nil
}
