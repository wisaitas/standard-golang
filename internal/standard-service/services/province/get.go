package province

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
	GetProvinces(query pkg.PaginationQuery) (resp []responses.ProvinceResponse, statusCode int, err error)
}

type get struct {
	provinceRepository repositories.ProvinceRepository
	redisUtil          pkg.RedisUtil
}

func NewGet(
	provinceRepository repositories.ProvinceRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		provinceRepository: provinceRepository,
		redisUtil:          redisUtil,
	}
}

func (r *get) GetProvinces(query pkg.PaginationQuery) (resp []responses.ProvinceResponse, statusCode int, err error) {
	provinces := []models.Province{}

	cacheKey := fmt.Sprintf("get_provinces:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return []responses.ProvinceResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []responses.ProvinceResponse{}, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.provinceRepository.GetAll(&provinces, &query, nil, nil); err != nil {
		return []responses.ProvinceResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, province := range provinces {
		respProvince := responses.ProvinceResponse{}
		resp = append(resp, respProvince.ModelToResponse(province))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return []responses.ProvinceResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return []responses.ProvinceResponse{}, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
