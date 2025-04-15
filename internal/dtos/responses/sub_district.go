package responses

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/internal/models"
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
	r.ID = model.ID
	r.CreatedAt = model.CreatedAt
	r.UpdatedAt = model.UpdatedAt
	r.CreatedBy = model.CreatedBy
	r.UpdatedBy = model.UpdatedBy
	r.NameTH = model.NameTh
	r.NameEN = model.NameEn
	r.DistrictID = model.DistrictID
	r.PostalCode = model.PostalCode

	return *r
}
