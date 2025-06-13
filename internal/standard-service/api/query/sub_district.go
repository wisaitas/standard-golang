package query

import (
	"github.com/google/uuid"
	"github.com/wisaitas/standard-golang/pkg"
)

type SubDistrictQuery struct {
	pkg.PaginationQuery
	DistrictID uuid.UUID `query:"district_id"`
}
