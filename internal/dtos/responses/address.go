package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type AddressResponse struct {
	pkg.BaseResponse
	ProvinceID    uuid.UUID `json:"province_id"`
	DistrictID    uuid.UUID `json:"district_id"`
	SubDistrictID uuid.UUID `json:"sub_district_id"`
	Address       string    `json:"address"`
}

func (r *AddressResponse) ModelToResponse(address models.Address) AddressResponse {
	r.ID = address.ID
	r.CreatedAt = address.CreatedAt
	r.UpdatedAt = address.UpdatedAt
	r.CreatedBy = address.CreatedBy
	r.UpdatedBy = address.UpdatedBy
	r.ProvinceID = address.ProvinceID
	r.DistrictID = address.DistrictID
	r.SubDistrictID = address.SubDistrictID
	r.Address = *address.Address

	return *r
}
