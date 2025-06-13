package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type AddressResponse struct {
	pkg.BaseResponse
	ProvinceID    uuid.UUID `json:"province_id"`
	DistrictID    uuid.UUID `json:"district_id"`
	SubDistrictID uuid.UUID `json:"sub_district_id"`
	Address       *string   `json:"address"`
}

func (r *AddressResponse) EntityToResponse(entity entity.Address) AddressResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.ProvinceID = entity.ProvinceID
	r.DistrictID = entity.DistrictID
	r.SubDistrictID = entity.SubDistrictID
	r.Address = entity.Address

	return *r
}
