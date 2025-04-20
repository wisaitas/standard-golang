package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/standard-service/models"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictResponse struct {
	pkg.BaseResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	DistrictID uuid.UUID `json:"district_id"`
	PostalCode string    `json:"postal_code"`
}

func (r *SubDistrictResponse) ModelToResponse(model models.SubDistrict) SubDistrictResponse {
	r.BaseResponse.ModelToResponse(model.BaseModel)
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.DistrictID = model.DistrictID
	r.PostalCode = model.PostalCode

	return *r
}
