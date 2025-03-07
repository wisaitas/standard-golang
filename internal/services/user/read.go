package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/wisaitas/standard-golang/internal/dtos/queries"
	"github.com/wisaitas/standard-golang/internal/dtos/responses"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/internal/repositories"
	"github.com/wisaitas/standard-golang/internal/utils"
)

type Read interface {
	GetUsers(query queries.PaginationQuery) (resp []responses.GetUsersResponse, statusCode int, err error)
}

type read struct {
	userRepository repositories.UserRepository
	redisUtil      utils.RedisClient
}

func NewRead(
	userRepository repositories.UserRepository,
	redisUtil utils.RedisClient,
) Read {
	return &read{
		userRepository: userRepository,
		redisUtil:      redisUtil,
	}
}

func (r *read) GetUsers(query queries.PaginationQuery) (resp []responses.GetUsersResponse, statusCode int, err error) {
	users := []models.User{}

	cacheKey := fmt.Sprintf("get_users:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []responses.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.userRepository.GetAll(&users, &query, nil); err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	for _, user := range users {
		respGetUser := responses.GetUsersResponse{}
		resp = append(resp, respGetUser.ModelToResponse(user))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []responses.GetUsersResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	return resp, http.StatusOK, nil
}
