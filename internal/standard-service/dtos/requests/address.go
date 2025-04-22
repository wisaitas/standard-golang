package requests

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
)

type AddressRequest struct {
	ProvinceID    uuid.UUID `json:"province_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174000"`
	DistrictID    uuid.UUID `json:"district_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174001"`
	SubDistrictID uuid.UUID `json:"sub_district_id" validate:"required" example:"123e4567-e89b-12d3-a456-426614174002"`
	Address       *string   `json:"address" example:"123 Main Street, Apartment 4B"`
}

func (r *AddressRequest) ReqToModel() models.Address {
	return models.Address{
		ProvinceID:    r.ProvinceID,
		DistrictID:    r.DistrictID,
		SubDistrictID: r.SubDistrictID,
		Address:       r.Address,
	}
}
