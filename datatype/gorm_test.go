package datatype

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// go test -run Test_GORM_GetSlicedItemsAndCount
func Test_GORM_GetSlicedItemsAndCount(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	g := GORM{}

	type User struct {
		ID   string
		Name string
	}
	err = db.AutoMigrate(&User{})
	assert.NoError(t, err)

	var createUsers []User
	for i := 1; i <= 500; i++ {
		createUsers = append(createUsers, User{
			ID:   strconv.Itoa(i),
			Name: "name " + strconv.Itoa(i),
		})
	}
	err = db.CreateInBatches(createUsers, 100).Error
	assert.NoError(t, err)

	_, totalSize, err := g.GetSlicedItemsAndCount(db.Model(&User{}), 0, 50, 50)
	assert.Equal(t, nil, err)
	assert.Equal(t, 500, totalSize)
}
