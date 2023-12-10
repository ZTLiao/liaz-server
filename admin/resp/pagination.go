package resp

type Pagination struct {
	PageNum  int         `json:"pageNum"`
	PageSize int         `json:"pageSize"`
	StartRow int         `json:"startRow"`
	EndRow   int         `json:"endRow"`
	Total    int64       `json:"total"`
	Pages    int         `json:"pages"`
	Records  interface{} `json:"records"`
}

func NewPagination(pageNum int, pageSize int) *Pagination {
	var pagination = new(Pagination)
	pagination.PageNum = pageNum
	pagination.PageSize = pageSize
	pagination.calculateStartAndEndRow()
	return pagination
}

func (e *Pagination) calculateStartAndEndRow() {
	if e.PageNum > 0 {
		e.StartRow = (e.PageNum - 1) * e.PageSize
		e.EndRow = e.StartRow + e.PageSize*1
	} else {
		e.StartRow = 0
		e.EndRow = e.StartRow
	}
}

func (e *Pagination) SetRecords(records interface{}) {
	e.Records = records
}

func (e *Pagination) SetTotal(total int64) {
	e.Total = total
	if e.Total == -1 {
		e.Pages = 1
		return
	}
	if e.PageSize > 0 {
		if total%int64(e.PageSize) == 0 {
			e.Pages = int(total) / e.PageSize
		} else {
			e.Pages = int(total)/e.PageSize + 1
		}
	} else {
		e.Pages = 0
	}
}
