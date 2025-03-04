package queries

type SubDistrictQuery struct {
	PaginationQuery
	DistrictID int `query:"district_id"`
}
