package initial

import "github.com/wisaitas/standard-golang/pkg"

type lib struct {
	redis     pkg.Redis
	jwt       pkg.JWT
	validator pkg.Validator
	bcrypt    pkg.Bcrypt
}

func newLib(clientConfig *clientConfig) *lib {
	return &lib{
		redis:     pkg.NewRedis(clientConfig.Redis),
		jwt:       pkg.NewJWT(),
		validator: pkg.NewValidator(),
		bcrypt:    pkg.NewBcrypt(),
	}
}
