package request

type PaginationParam struct {
	Page     *int    `query:"page"`
	PageSize *int    `query:"page_size"`
	Sort     *string `query:"sort"`
	Order    *string `query:"order"`
}
