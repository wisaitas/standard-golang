package district

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
	"github.com/wisaitas/standard-golang/internal/standard-service/api/query"
	"github.com/wisaitas/standard-golang/internal/standard-service/api/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/internal/standard-service/repository"
)

type Get interface {
	GetDistricts(query query.DistrictQuery) (resp []response.DistrictResponse, statusCode int, err error)
}

type get struct {
	districtRepository repository.DistrictRepository
	redisUtil          redisPkg.Redis
}

func NewGet(
	districtRepository repository.DistrictRepository,
	redisUtil redisPkg.Redis,
) Get {
	return &get{
		districtRepository: districtRepository,
		redisUtil:          redisUtil,
	}
}

func (r *get) GetDistricts(query query.DistrictQuery) (resp []response.DistrictResponse, statusCode int, err error) {
	districts := []entity.District{}

	cacheKey := fmt.Sprintf("get_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.ProvinceID)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redisLib.Nil {
		return nil, http.StatusInternalServerError, utils.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return nil, http.StatusInternalServerError, utils.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.districtRepository.GetAll(&districts, &query.PaginationQuery, repositoryPkg.NewCondition("province_id = ?", query.ProvinceID), nil); err != nil {
		return nil, http.StatusInternalServerError, utils.Error(err)
	}

	for _, district := range districts {
		respDistrict := response.DistrictResponse{}
		resp = append(resp, respDistrict.EntityToResponse(district))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, utils.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return nil, http.StatusInternalServerError, utils.Error(err)
	}

	return resp, http.StatusOK, nil
}
