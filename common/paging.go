package common

type Paging struct {
	PageNumber int32 `json:"page_number,omitempty" form:"page_number" binding:"omitempty,min=1"`
	PageSize   int32 `json:"page_size,omitempty" form:"page_size" binding:"omitempty,min=5,max=50"`
	Total      int64 `json:"total,omitempty" form:"total" binding:"omitempty"`
}

func (p *Paging) FillDefault() {
	if p.PageNumber == 0 {
		p.PageNumber = 1
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
