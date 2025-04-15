package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	ProvinceID uuid.UUID `json:"province_id"`
}

func (r *DistrictResponse) ModelToResponse(model models.District) DistrictResponse {
	r.ID = model.ID
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	r.CreatedBy = model.CreatedBy
	r.UpdatedBy = model.UpdatedBy
	r.NameTH = model.NameTh
	r.NameEN = model.NameEn
	r.ProvinceID = model.ProvinceID

	return *r
}
