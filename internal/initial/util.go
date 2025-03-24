package initial

import "github.com/wisaitas/standard-golang/pkg"

type Util struct {
	RedisUtil       pkg.RedisUtil
	JWTUtil         pkg.JWTUtil
	TransactionUtil pkg.TransactionUtil
	ValidatorUtil   pkg.ValidatorUtil
	BcryptUtil      pkg.BcryptUtil
}

func NewUtil(configs *Configs) *Util {
	return &Util{
		RedisUtil:       pkg.NewRedisUtil(configs.Redis),
		JWTUtil:         pkg.NewJWTUtil(),
		TransactionUtil: pkg.NewTransactionUtil(configs.DB),
		ValidatorUtil:   pkg.NewValidatorUtil(),
		BcryptUtil:      pkg.NewBcrypt(),
	}
}
