package paginations

type Pagination struct {
	TotalCount int
	// 页码, 从1开始
	PageNum  int
	PageSize int
}

func NewPagination(totalCount int, pageNum int, pageSize int) *Pagination {
	return &Pagination{
		TotalCount: totalCount,
		PageNum:    pageNum,
		PageSize:   pageSize,
	}
}

func (p *Pagination) TotalPage() int {
	totalPage := 0
	if p.TotalCount%p.PageSize == 0 {
		totalPage = p.TotalCount / p.PageSize
	} else {
		totalPage = p.TotalCount/p.PageSize + 1
	}
	return totalPage
}

func (p *Pagination) HasNext() bool {
	return p.PageNum < p.TotalPage()
}

func (p *Pagination) HasPrev() bool {
	return p.PageNum > 1
}

func (p *Pagination) Offset() int {
	return (p.PageNum - 1) * p.PageSize
}

func (p *Pagination) Limit() int {
	return p.PageSize
}
