package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type Get interface {
	GetUsers(query repositoryPkg.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error)
}

type get struct {
	userRepository repository.UserRepository
	redisUtil      redisPkg.Redis
}

func NewGet(
	userRepository repository.UserRepository,
	redisUtil redisPkg.Redis,
) Get {
	return &get{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *get) GetUsers(query repositoryPkg.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error) {
	users := []entity.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.userRepository.GetAll(&users, &query, nil, nil); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	for _, user := range users {
		respGetUser := response.GetUsersResponse{}
		resp = append(resp, respGetUser.EntityToResponse(user))
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJSON, 10*time.Second); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	return resp, http.StatusOK, nil
}
