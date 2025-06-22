package initial

import (
	"github.com/wisaitas/share-pkg/auth/jwt"
	"github.com/wisaitas/share-pkg/cache/redis"
	"github.com/wisaitas/share-pkg/crypto/bcrypt"
	transactionmanager "github.com/wisaitas/share-pkg/db/transaction-manager"
	"github.com/wisaitas/share-pkg/validator"
)

type sharePkg struct {
	redis              redis.Redis
	jwt                jwt.Jwt
	validator          validator.Validator
	bcrypt             bcrypt.Bcrypt
	transactionManager transactionmanager.TransactionManager
}

func newSharePkg(clientConfig *clientConfig) *sharePkg {
	return &sharePkg{
		redis:              redis.NewRedis(clientConfig.Redis),
		jwt:                jwt.NewJwt(),
		validator:          validator.NewValidator(),
		bcrypt:             bcrypt.NewBcrypt(),
		transactionManager: transactionmanager.NewTransactionManager(clientConfig.DB),
	}
}
