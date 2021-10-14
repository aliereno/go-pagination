package datatype

type IDatatype interface {
	GetSlicedItemsAndCount(data interface{}, start, end, pageSize int) ([]interface{}, int, error)
}
