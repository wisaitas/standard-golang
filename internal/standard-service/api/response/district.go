package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type DistrictResponse struct {
	response.EntityResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	ProvinceID uuid.UUID `json:"province_id"`
}

func (r *DistrictResponse) EntityToResponse(entity entity.District) DistrictResponse {
	r.EntityResponse = entity.EntityToResponse()
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN
	r.ProvinceID = entity.ProvinceID

	return *r
}
