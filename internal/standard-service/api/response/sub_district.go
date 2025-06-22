package response

import (
	"github.com/google/uuid"
	"github.com/wisaitas/share-pkg/response"
	"github.com/wisaitas/standard-golang/internal/standard-service/entity"
)

type SubDistrictResponse struct {
	response.EntityResponse
	NameTH     string    `json:"name_th"`
	NameEN     string    `json:"name_en"`
	DistrictID uuid.UUID `json:"district_id"`
	PostalCode string    `json:"postal_code"`
}

func (r *SubDistrictResponse) EntityToResponse(entity entity.SubDistrict) SubDistrictResponse {
	r.EntityResponse = entity.EntityToResponse()
	r.NameTH = entity.NameTH
	r.NameEN = entity.NameEN
	r.DistrictID = entity.DistrictID
	r.PostalCode = entity.PostalCode

	return *r
}
