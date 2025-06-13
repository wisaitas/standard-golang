package response

import (
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
	"github.com/wisaitas/standard-golang/pkg"
)

type ProvinceResponse struct {
	pkg.BaseResponse
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *ProvinceResponse) EntityToResponse(entity entity.Province) ProvinceResponse {
	r.BaseResponse.EntityToResponse(entity.BaseEntity)
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN

	return *r
}
