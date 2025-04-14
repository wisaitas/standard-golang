package queries

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type DistrictQuery struct {
	pkg.PaginationQuery
	ProvinceID uuid.UUID `query:"province_id"`
}
