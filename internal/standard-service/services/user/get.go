package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/standard-service/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/internal/standard-service/repositories"
	"github.com/wisaitas/standard-golang/pkg"
)

type Get interface {
	GetUsers(query pkg.PaginationQuery) (resp []responses.GetUsersResponse, statusCode int, err error)
}

type get struct {
	userRepository repositories.UserRepository
	redisUtil      pkg.RedisUtil
}

func NewGet(
	userRepository repositories.UserRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *get) GetUsers(query pkg.PaginationQuery) (resp []responses.GetUsersResponse, statusCode int, err error) {
	users := []models.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []responses.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.userRepository.GetAll(&users, &query, nil, nil); err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, user := range users {
		respGetUser := responses.GetUsersResponse{}
		resp = append(resp, respGetUser.ModelToResponse(user))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
