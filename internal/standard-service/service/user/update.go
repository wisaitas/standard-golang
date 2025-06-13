package user

import (
	"net/http"

	"github.com/wisaitas/standard-golang/internal/standard-service/api/param"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/constants"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"github.com/wisaitas/standard-golang/pkg"
)

type Update interface {
	UpdateUser(param param.UserParam, req request.UpdateUserRequest) (resp response.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository        repository.UserRepository
	userHistoryRepository repository.UserHistoryRepository
	redis                 pkg.Redis
}

func NewUpdate(
	userRepository repository.UserRepository,
	userHistoryRepository repository.UserHistoryRepository,
	redis pkg.Redis,
) Update {
	return &update{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redis:                 redis,
	}
}

func (r *update) UpdateUser(param param.UserParam, request request.UpdateUserRequest) (resp response.UpdateUserResponse, statusCode int, err error) {
	user := entity.User{}

	relations := []pkg.Relation{
		{
			Query: "Addresses",
		},
	}

	if err := r.userRepository.GetBy(&user, pkg.NewCondition("id = ?", param.ID), &relations); err != nil {
		return resp, http.StatusNotFound, pkg.Error(err)
	}

	tm := pkg.NewTxManager(r.userRepository.GetDB())

	txUserRepository := r.userRepository.WithTxManager(tm)
	txUserHistoryRepository := r.userHistoryRepository.WithTxManager(tm)

	userBeforeUpdate := entity.UserHistory{
		Action:       constants.Action.Update,
		OldFirstName: user.FirstName,
		OldLastName:  user.LastName,
		OldBirthDate: user.BirthDate,
		OldPassword:  user.Password,
		OldEmail:     user.Email,
		OldVersion:   user.Version,
	}

	if err := txUserHistoryRepository.Create(&userBeforeUpdate); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
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
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := tm.Commit(); err != nil {
		return resp, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp.EntityToResponse(user), http.StatusOK, nil
}
