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

type Get interface {
	GetSubDistricts(query queries.SubDistrictQuery) (resp []responses.SubDistrictResponse, statusCode int, err error)
}

type get struct {
	subDistrictRepository repositories.SubDistrictRepository
	redisUtil             pkg.RedisUtil
}

func NewGet(
	subDistrictRepository repositories.SubDistrictRepository,
	redisUtil pkg.RedisUtil,
) Get {
	return &get{
		subDistrictRepository: subDistrictRepository,
		redisUtil:             redisUtil,
	}
}

func (r *get) GetSubDistricts(query queries.SubDistrictQuery) (resp []responses.SubDistrictResponse, statusCode int, err error) {
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

	if err := r.subDistrictRepository.GetAll(&subDistricts, &query.PaginationQuery, pkg.NewCondition("district_id = ?", query.DistrictID), nil); err != nil {
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
