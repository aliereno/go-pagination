package datatype

type IDatabase interface {
	GetSlicedItemsAndCount(data interface{}, start, end, pageSize int) ([]interface{}, int, error)
}
