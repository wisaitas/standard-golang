package pkg

import "golang.org/x/crypto/bcrypt"

type BcryptUtil interface {
	GenerateFromPassword(password string, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}

type bcryptUtil struct {
}

func NewBcrypt() BcryptUtil {
	return &bcryptUtil{}
}

func (r *bcryptUtil) GenerateFromPassword(password string, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), cost)
}

func (r *bcryptUtil) CompareHashAndPassword(hashedPassword, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}
