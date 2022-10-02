package common

type Paging struct {
	PageNumber int32 `json:"page_number,omitempty" form:"page_number" binding:"omitempty,min=1"`
	PageSize   int32 `json:"page_size,omitempty" form:"page_size" binding:"omitempty,min=5,max=50"`
	Total      int64 `json:"total,omitempty" form:"total" binding:"omitempty"`
}

func (paging *Paging) FillDefault() {
	if paging.PageNumber == 0 {
		paging.PageNumber = 1
	}

	if paging.PageSize == 0 {
		paging.PageSize = 10
	}
}
