package responses

import "github.com/wisaitas/standard-golang/internal/models"

type ProvinceResponse struct {
	ID     int    `json:"id"`
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *ProvinceResponse) ModelToResponse(model models.Province) ProvinceResponse {
	r.ID = model.ID
	r.NameTH = model.NameTH
	r.NameEN = model.NameEN

	return *r
}
