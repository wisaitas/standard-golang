package responses

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type ProvinceResponse struct {
	pkg.BaseResponse
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *ProvinceResponse) ModelToResponse(model models.Province) ProvinceResponse {
	r.ID = model.ID
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	r.CreatedBy = model.CreatedBy
	r.UpdatedBy = model.UpdatedBy
	r.NameTH = model.NameTh
	r.NameEN = model.NameEn

	return *r
}
