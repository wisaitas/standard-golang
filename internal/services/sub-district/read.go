package sub_district

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
	GetSubDistricts(query queries.SubDistrictQuery) (resp []responses.SubDistrictResponse, statusCode int, err error)
}

type read struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisUtil             pkg.RedisClient
}

func NewRead(
	subDistrictRepository repositories.SubDistrictRepository,
	redisUtil pkg.RedisClient,
) Read {
	return &read{
		subDistrictRepository: subDistrictRepository,
		redisUtil:             redisUtil,
	}
}

func (r *read) GetSubDistricts(query queries.SubDistrictQuery) (resp []responses.SubDistrictResponse, statusCode int, err error) {
	subDistricts := []models.SubDistrict{}

	cacheKey := fmt.Sprintf("get_sub_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.DistrictID)

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

	if err := r.subDistrictRepository.GetAll(&subDistricts, &query.PaginationQuery, map[string]interface{}{"district_id": query.DistrictID}); err != nil {
		return nil, http.StatusInternalServerError, pkg.Error(err)
	}

	for _, subDistrict := range subDistricts {
		respSubDistrict := responses.SubDistrictResponse{}
		resp = append(resp, respSubDistrict.ModelToResponse(subDistrict))
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
