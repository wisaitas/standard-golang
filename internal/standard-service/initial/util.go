package initial

import "github.com/wisaitas/standard-golang/pkg"

type util struct {
	redisUtil       pkg.RedisUtil
	jwtUtil         pkg.JWTUtil
	transactionUtil pkg.TransactionUtil
	validatorUtil   pkg.ValidatorUtil
	bcryptUtil      pkg.BcryptUtil
}

func newUtil(clientConfig *clientConfig) *util {
	return &util{
		redisUtil:       pkg.NewRedisUtil(clientConfig.Redis),
		jwtUtil:         pkg.NewJWTUtil(),
		transactionUtil: pkg.NewTransactionUtil(clientConfig.DB),
		validatorUtil:   pkg.NewValidatorUtil(),
		bcryptUtil:      pkg.NewBcrypt(),
	}
}
