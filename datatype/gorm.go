package datatype

import (
	"errors"
	"reflect"

	"gorm.io/gorm"
)

type GORM struct {
}

func (g GORM) GetSlicedItemsAndCount(data interface{}, start, end, pageSize int) (interface{}, int, error) {
	query, ok := data.(*gorm.DB)
	if !ok {
		return nil, 0, errors.New("query is not valid")
	}
	var totalRows int64
	err := query.Count(&totalRows).Error
	if err != nil {
		return nil, 0, err
	}
	totalSize := int(totalRows)
	dtype := reflect.TypeOf(query.Statement.Model)
	pages := reflect.New(reflect.SliceOf(dtype)).Interface()
	query.Offset(start).Limit(pageSize).Find(pages)
	return pages, totalSize, nil
}
