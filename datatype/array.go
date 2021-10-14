package datatype

import (
	"errors"
	"reflect"
)

type Array struct {
}

func (a Array) GetSlicedItemsAndCount(data interface{}, start, end, pageSize int) ([]interface{}, int, error) {
	// Check slice is valid
	slice, valid := takeArg(data, reflect.Slice)
	if !valid {
		return nil, 0, errors.New("slice is not valid")
	}

	currentPageSize := pageSize
	if end > slice.Len() {
		currentPageSize = currentPageSize - (end - slice.Len())
	}
	if currentPageSize < 0 {
		currentPageSize = 0
	}
	items := make([]interface{}, currentPageSize)
	for i := 0; i < currentPageSize; i++ {
		if (i + start) < slice.Len() {
			items[i] = slice.Index(i + start).Interface()
		}
	}
	return items, slice.Len(), nil
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}
