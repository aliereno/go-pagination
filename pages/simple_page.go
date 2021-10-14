package pages

type SimplePage struct {
	Page      int           `json:"currentPage"`
	PageSize  int           `json:"pageSize"`
	TotalPage int           `json:"totalPage"`
	Items     []interface{} `json:"items"`
}

func (s SimplePage) Response(page int, pageSize int, totalPage int, items []interface{}, query string) interface{} {
	res := new(SimplePage)
	res.Page = page
	res.PageSize = pageSize
	res.TotalPage = totalPage
	res.Items = items
	return res
}
