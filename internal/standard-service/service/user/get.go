package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
	"github.com/wisaitas/standard-golang/pkg"
)

type Get interface {
	GetUsers(query pkg.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error)
}

type get struct {
	userRepository repository.UserRepository
	redisUtil      pkg.Redis
}

func NewGet(
	userRepository repository.UserRepository,
	redisUtil pkg.Redis,
) Get {
	return &get{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *get) GetUsers(query pkg.PaginationQuery) (resp []response.GetUsersResponse, statusCode int, err error) {
	users := []entity.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.userRepository.GetAll(&users, &query, nil, nil); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, user := range users {
		respGetUser := response.GetUsersResponse{}
		resp = append(resp, respGetUser.EntityToResponse(user))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []response.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
