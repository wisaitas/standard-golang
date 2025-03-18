package queries

import "github.com/wisaitas/standard-golang/pkg"

type DistrictQuery struct {
	pkg.PaginationQuery
	ProvinceID int `query:"province_id"`
}
