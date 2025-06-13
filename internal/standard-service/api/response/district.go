package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	ProvinceID uuid.UUID `json:"province_id"`
}

func (r *DistrictResponse) EntityToResponse(entity entity.District) DistrictResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN
	r.ProvinceID = entity.ProvinceID

	return *r
}
