package responses

import "github.com/wisaitas/standard-golang/internal/models"

type SubDistrictResponse struct {
	ID         int    `json:"id"`
	NameTH     string `json:"name_th"`
	NameEN     string `json:"name_en"`
	DistrictID int    `json:"district_id"`
	ZipCode    int    `json:"zip_code"`
}

func (r *SubDistrictResponse) ModelToResponse(model models.SubDistrict) SubDistrictResponse {
	r.ID = model.ID
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN
	r.DistrictID = model.DistrictID
	r.ZipCode = model.ZipCode

	return *r
}
