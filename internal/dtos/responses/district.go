package responses

import "github.com/wisaitas/standard-golang/internal/models"

type DistrictResponse struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	ProvinceID int    `json:"province_id"`
}

func (r *DistrictResponse) ModelToResponse(model models.District) DistrictResponse {
	r.ID = model.ID
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.ProvinceID = model.ProvinceID

	return *r
}
