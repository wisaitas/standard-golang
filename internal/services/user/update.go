package user

import (
	"net/http"

	"github.com/wisaitas/standard-golang/internal/constants"
	"github.com/wisaitas/standard-golang/internal/dtos/params"
	"github.com/wisaitas/standard-golang/internal/dtos/requests"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type Update interface {
	UpdateUser(param params.UserParams, req requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository        repositories.UserRepository
	userHistoryRepository repositories.UserHistoryRepository
	redisUtil             utils.RedisClient
}

func NewUpdate(
	userRepository repositories.UserRepository,
	userHistoryRepository repositories.UserHistoryRepository,
	redisUtil utils.RedisClient,
) Update {
	return &update{
		userRepository:        userRepository,
		userHistoryRepository: userHistoryRepository,
		redisUtil:             redisUtil,
	}
}

func (r *update) UpdateUser(param params.UserParams, request requests.UpdateUserRequest) (resp responses.UpdateUserResponse, statusCode int, err error) {
	user := models.User{}

	if err := r.userRepository.GetBy(map[string]any{"id": param.ID}, &user, "Addresses"); err != nil {
		return resp, 0, err
	}

	tx := r.userRepository.BeginTx()

	if err := tx.Create(
		&models.UserHistory{
			Action:    constants.ACTION.UPDATE,
			UserID:    user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			BirthDate: user.BirthDate,
			Version:   user.Version,
		},
	).Error; err != nil {
		tx.Rollback()
		return resp, 0, err
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

	if err := tx.Updates(&user).Error; err != nil {
		tx.Rollback()
		return resp, 0, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return resp, 0, err
	}

	respInfo := responses.UpdateUserResponse{}
	resp = respInfo.ModelToResponse(user)

	return resp, http.StatusOK, nil
}
