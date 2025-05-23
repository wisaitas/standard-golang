package user

import (
	"net/http"

	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/params"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
	"gorm.io/gorm"
)

type Update interface {
	UpdateUser(param params.UserParams, req requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	transactionUtil       pkg.TransactionUtil
	redisUtil             pkg.RedisUtil
}

func NewUpdate(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	transactionUtil pkg.TransactionUtil,
	redisUtil pkg.RedisUtil,
) Update {
	return &update{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		transactionUtil:       transactionUtil,
		redisUtil:             redisUtil,
	}
}

func (r *update) UpdateUser(param params.UserParams, request requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error) {
	user := models.User{}

	relations := []pkg.Relation{
		{
			Query: "Addresses",
		},
	}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("id = ?", param.ID), &relations); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	if err := r.transactionUtil.ExecuteInTransaction(func(tx *gorm.DB) error {
		txUserRepository := r.userRepository.WithTx(tx)
		txUserHistoryRepository := r.userHistoryRepository.WithTx(tx)

		userBeforeUpdate := models.UserHistory{
			Action:       constants.Action.Update,
			OldFirstName: user.FirstName,
			OldLastName:  user.LastName,
			OldBirthDate: user.BirthDate,
			OldPassword:  user.Password,
			OldEmail:     user.Email,
			OldVersion:   user.Version,
		}

		if err := txUserHistoryRepository.Create(&userBeforeUpdate); err != nil {
			return pkg.Error(err)
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
			return pkg.Error(err)
		}

		return nil
	}); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.ModelToResponse(user), http.StatusOK, nil
}
