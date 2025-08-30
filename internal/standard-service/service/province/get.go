package province

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	redisLib "github.com/redis/go-redis/v9"
	redisPkg "github.com/wisaitas/share-pkg/cache/redis"
	repositoryPkg "github.com/wisaitas/share-pkg/db/repository"
	"github.com/wisaitas/share-pkg/utils"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type Get interface {
	GetProvinces(query repositoryPkg.PaginationQuery) (resp []response.ProvinceResponse, statusCode int, err error)
}

type get struct {
	provinceRepository repository.ProvinceRepository
	redisUtil          redisPkg.Redis
}

func NewGet(
	provinceRepository repository.ProvinceRepository,
	redisUtil redisPkg.Redis,
) Get {
	return &get{
		provinceRepository: provinceRepository,
		redisUtil:          redisUtil,
	}
}

func (r *get) GetProvinces(query repositoryPkg.PaginationQuery) (resp []response.ProvinceResponse, statusCode int, err error) {
	provinces := []entity.Province{}

	cacheKey := fmt.Sprintf("get_provinces:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redisLib.Nil {
		return []response.ProvinceResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return []response.ProvinceResponse{}, http.StatusInternalServerError, utils.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.provinceRepository.GetAll(&provinces, &query, nil, nil); err != nil {
		return []response.ProvinceResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	for _, province := range provinces {
		respProvince := response.ProvinceResponse{}
		resp = append(resp, respProvince.EntityToResponse(province))
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		return []response.ProvinceResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJSON, 10*time.Second); err != nil {
		return []response.ProvinceResponse{}, http.StatusInternalServerError, utils.Error(err)
	}

	return resp, http.StatusOK, nil
}
