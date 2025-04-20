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
	r.BaseResponse.ModelToResponse(model.BaseModel)
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN

	return *r
}
