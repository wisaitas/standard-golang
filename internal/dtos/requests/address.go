package requests

import "github.com/wisaitas/standard-golang/internal/models"

type AddressRequest struct {
	ProvinceID    int     `json:"province_id" validate:"required"`
	DistrictID    int     `json:"district_id" validate:"required"`
	SubDistrictID int     `json:"sub_district_id" validate:"required"`
	Address       *string `json:"address"`
}

func (r *AddressRequest) ReqToModel() models.Address {
	return models.Address{
		ProvinceID:    r.ProvinceID,
		DistrictID:    r.DistrictID,
		SubDistrictID: r.SubDistrictID,
		Address:       r.Address,
	}
}
