package district

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
	"github.com/wisaitas/standard-golang/pkg"
)

type Read interface {
	GetDistricts(query queries.DistrictQuery) (resp []responses.DistrictResponse, statusCode int, err error)
}

type read struct {
	districtRepository repositories.DistrictRepository
	redisUtil          pkg.RedisUtil
}

func NewRead(
	districtRepository repositories.DistrictRepository,
	redisUtil pkg.RedisUtil,
) Read {
	return &read{
		districtRepository: districtRepository,
		redisUtil:          redisUtil,
	}
}

func (r *read) GetDistricts(query queries.DistrictQuery) (resp []responses.DistrictResponse, statusCode int, err error) {
	districts := []models.District{}

	cacheKey := fmt.Sprintf("get_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.ProvinceID)

	cache, err := r.redisUtil.Get(context.Background(), cacheKey)
	if err != nil && err != redis.Nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	if cache != "" {
		if err := json.Unmarshal([]byte(cache), &resp); err != nil {
			return nil, http.StatusInternalServerError, pkg.Error(err)
		}

		return resp, http.StatusOK, nil
	}

	if err := r.districtRepository.GetAll(&districts, &query.PaginationQuery, map[string]interface{}{"province_id": query.ProvinceID}); err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, district := range districts {
		respDistrict := responses.DistrictResponse{}
		resp = append(resp, respDistrict.ModelToResponse(district))
	}

	respJson, err := json.Marshal(resp)
	if err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	if err := r.redisUtil.Set(context.Background(), cacheKey, respJson, 10*time.Second); err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	return resp, http.StatusOK, nil
}
