package response

import (
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type ProvinceResponse struct {
	response.EntityResponse
	NameTH string `json:"name_th"`
	NameEN string `json:"name_en"`
}

func (r *ProvinceResponse) EntityToResponse(entity entity.Province) ProvinceResponse {
	r.EntityResponse = entity.EntityToResponse()
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN

	return *r
}
