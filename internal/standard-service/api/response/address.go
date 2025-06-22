package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type AddressResponse struct {
	response.EntityResponse
	ProvinceID    uuid.UUID `json:"province_id"`
	DistrictID    uuid.UUID `json:"district_id"`
	SubDistrictID uuid.UUID `json:"sub_district_id"`
	Address       *string   `json:"address"`
}

func (r *AddressResponse) EntityToResponse(entity entity.Address) AddressResponse {
	r.EntityResponse = entity.EntityToResponse()
	r.ProvinceID = entity.ProvinceID
	r.DistrictID = entity.DistrictID
	r.SubDistrictID = entity.SubDistrictID
	r.Address = entity.Address

	return *r
}
