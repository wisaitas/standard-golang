package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	DistrictID uuid.UUID `json:"district_id"`
	PostalCode string    `json:"postal_code"`
}

func (r *SubDistrictResponse) EntityToResponse(entity entity.SubDistrict) SubDistrictResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN
	r.DistrictID = entity.DistrictID
	r.PostalCode = entity.PostalCode

	return *r
}
