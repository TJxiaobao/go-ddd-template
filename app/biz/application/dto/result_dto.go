package dto

type PageResult struct {
	PageInfo PageInfo    `json:"page_info"`
	Rows     interface{} `json:"rows"`
}

type ListResult struct {
	TotalCount int         `json:"total_count"`
	Rows       interface{} `json:"rows"`
}

type PageInfo struct {
	TotalNum    int64 `json:"total_num"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
}

type ReqPageInfo struct {
	PageNum  int `form:"page_num,default=1"    binding:"omitempty,min=1" json:"page_num"`
	PageSize int `form:"page_size,default=10"  binding:"omitempty,min=1,max=5000" json:"page_size"`
}

func NewPageResult(rows interface{}, pageNum, pageSize int) *PageResult {
	return &PageResult{
		PageInfo: PageInfo{
			PageSize:    pageSize,
			CurrentPage: pageNum,
		},
		Rows: rows,
	}
}
