package initial

import (
	"github.com/wisaitas/standard-golang/internal/validates"
)

func InitializeValidates() *Validates {
	return &Validates{
		UserValidate: *validates.NewUserValidate(),
		AuthValidate: *validates.NewAuthValidate(),
	}
}

type Validates struct {
	UserValidate validates.UserValidate
	AuthValidate validates.AuthValidate
}
