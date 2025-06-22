package sub_district

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
	GetSubDistricts(query query.SubDistrictQuery) (resp []response.SubDistrictResponse, statusCode int, err error)
}

type get struct {
	subDistrictRepository repository.SubDistrictRepository
	redisUtil             redisPkg.Redis
}

func NewGet(
	subDistrictRepository repository.SubDistrictRepository,
	redisUtil redisPkg.Redis,
) Get {
	return &get{
		subDistrictRepository: subDistrictRepository,
		redisUtil:             redisUtil,
	}
}

func (r *get) GetSubDistricts(query query.SubDistrictQuery) (resp []response.SubDistrictResponse, statusCode int, err error) {
	subDistricts := []entity.SubDistrict{}

	cacheKey := fmt.Sprintf("get_sub_districts:%v:%v:%v:%v:%v", query.Page, query.PageSize, query.Sort, query.Order, query.DistrictID)

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

	if err := r.subDistrictRepository.GetAll(&subDistricts, &query.PaginationQuery, repositoryPkg.NewCondition("district_id = ?", query.DistrictID), nil); err != nil {
		return nil, http.StatusInternalServerError, utils.Error(err)
	}

	for _, subDistrict := range subDistricts {
		respSubDistrict := response.SubDistrictResponse{}
		resp = append(resp, respSubDistrict.EntityToResponse(subDistrict))
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
