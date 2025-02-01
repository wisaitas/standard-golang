package request

type PaginationParam struct {
	Page     *int    `json:"page" form:"page"`
	PageSize *int    `json:"page_size" form:"page_size"`
	Sort     *string `json:"sort" form:"sort"`
	Order    *string `json:"order" form:"order"`
}
