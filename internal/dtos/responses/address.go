package responses

import (
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type AddressResponse struct {
	pkg.BaseResponse
	ProvinceID    int    `json:"province_id"`
	DistrictID    int    `json:"district_id"`
	SubDistrictID int    `json:"sub_district_id"`
	Address       string `json:"address"`
}

func (r *AddressResponse) ModelToResponse(address models.Address) AddressResponse {
	r.ID = address.ID
	r.CreatedAt = address.CreatedAt
	r.UpdatedAt = address.UpdatedAt
	r.ProvinceID = address.ProvinceID
	r.DistrictID = address.DistrictID
	r.SubDistrictID = address.SubDistrictID
	r.Address = *address.Address

	return *r
}
