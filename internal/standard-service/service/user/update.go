package user

import (
	"net/http"

	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/param"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/request"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type Update interface {
	UpdateUser(param param.UserParam, req request.UpdateUserRequest) (resp response.UpdateUserResponse, statusCode int, err error)
}

type update struct {
	userRepository repository.UserRepository
	redis          redisPkg.Redis
}

func NewUpdate(
	userRepository repository.UserRepository,
	redis redisPkg.Redis,
) Update {
	return &update{
		userRepository: userRepository,
		redis:          redis,
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

	if err := r.userRepository.Update(&user); err != nil {
		return resp, http.StatusInternalServerError, utils.Error(err)
	}

	return resp.EntityToResponse(user), http.StatusOK, nil
}
