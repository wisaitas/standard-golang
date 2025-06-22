package user

import (
	"net/http"

	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	transactionmanager "github.com/wisaitas/share-pkg/db/transaction-manager"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/param"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/constant"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"gorm.io/gorm"
)

type Update interface {
	UpdateUser(param param.UserParam, req request.UpdateUserRequest) (resp response.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository        repository.UserRepository
	userHistoryRepository repository.UserHistoryRepository
	redis                 redisPkg.Redis
	transactionManager    transactionmanager.TransactionManager
}

func NewUpdate(
	userRepository repository.UserRepository,
	userHistoryRepository repository.UserHistoryRepository,
	redis redisPkg.Redis,
	transactionManager transactionmanager.TransactionManager,
) Update {
	return &update{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redis:                 redis,
		transactionManager:    transactionManager,
	}
}

func (r *update) UpdateUser(param param.UserParam, request request.UpdateUserRequest) (resp response.UpdateUserResponse, statusCode int, err error) {
	user := entity.User{}

	relations := []repositoryPkg.Relation{
		{
			Query: "Addresses",
		},
	}

	if err := r.userRepository.GetBy(&user, repositoryPkg.NewCondition("id = ?", param.ID), &relations); err != nil {
		return resp, http.StatusNotFound, utils.Error(err)
	}

	if err := r.transactionManager.ExecuteInTransaction(func(tx *gorm.DB) error {
		txUserRepository := r.userRepository.WithTx(tx)
		txUserHistoryRepository := r.userHistoryRepository.WithTx(tx)

		userBeforeUpdate := entity.UserHistory{
			Action:       constant.Action.Update,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			OldBirthDate: user.BirthDate,
			OldPassword:  user.Password,
			OldEmail:     user.Email,
			OldVersion:   user.Version,
		}

		if err := txUserHistoryRepository.Create(&userBeforeUpdate); err != nil {
			return err
		}

		if request.FirstName != nil {
			user.FirstName = *request.FirstName
		}

		if request.LastName != nil {
			user.LastName = *request.LastName
		}

		if request.BirthDate != nil {
			user.BirthDate = *request.BirthDate
		}

		if request.Email != nil {
			user.Email = *request.Email
		}

		if err := txUserRepository.Update(&user); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	return resp.EntityToResponse(user), http.StatusOK, nil
}
