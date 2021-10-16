package pages

// IPage is used as response interface.
type IPage interface {
	Response(page int, pageSize int, totalPage int, items interface{}, query string) interface{}
}
