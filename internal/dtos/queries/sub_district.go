package queries

import "github.com/wisaitas/standard-golang/pkg"

type SubDistrictQuery struct {
	pkg.PaginationQuery
	DistrictID int `query:"district_id"`
}
