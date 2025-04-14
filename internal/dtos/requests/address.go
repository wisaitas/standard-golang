package requests

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/models"
)

type AddressRequest struct {
	ProvinceID    uuid.UUID `json:"province_id" validate:"required"`
	DistrictID    uuid.UUID `json:"district_id" validate:"required"`
	SubDistrictID uuid.UUID `json:"sub_district_id" validate:"required"`
	Address       *string   `json:"address"`
}

func (r *AddressRequest) ReqToModel() models.Address {
	return models.Address{
		ProvinceID:    r.ProvinceID,
		DistrictID:    r.DistrictID,
		SubDistrictID: r.SubDistrictID,
		Address:       r.Address,
	}
}
