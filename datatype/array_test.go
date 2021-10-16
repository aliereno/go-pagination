package datatype

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -run Test_Array_GetSlicedItemsAndCount
func Test_Array_GetSlicedItemsAndCount(t *testing.T) {
	arr := Array{}
	var items []interface{}
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			items = append(items, i)
		} else {
			items = append(items, strconv.Itoa(i))
		}
	}
	paginatedItems, totalSize, err := arr.GetSlicedItemsAndCount(items, 0, 50, 50)
	assert.Equal(t, nil, err)
	assert.Equal(t, 100, totalSize)
	assert.Equal(t, items[0:50], paginatedItems)
}
