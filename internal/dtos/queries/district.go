package queries

type DistrictQuery struct {
	PaginationQuery
	ProvinceID int `query:"province_id"`
}
