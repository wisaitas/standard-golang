package pkg

import (
	bcryptLib "golang.org/x/crypto/bcrypt"
)

type Bcrypt interface {
	GenerateFromPassword(password string, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type bcrypt struct {
}

func NewBcrypt() Bcrypt {
	return &bcrypt{}
}

func (r *bcrypt) GenerateFromPassword(password string, cost int) ([]byte, error) {
	return bcryptLib.GenerateFromPassword([]byte(password), cost)
}

func (r *bcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcryptLib.CompareHashAndPassword(hashedPassword, password)
}
