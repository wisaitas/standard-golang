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
	r.BaseResponse.ModelToResponse(model.BaseModel)
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.ProvinceID = model.ProvinceID

	return *r
}
