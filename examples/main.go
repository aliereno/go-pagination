package main

import (
	"log"
	"strconv"

	"github.com/aliereno/go-pagination"
	"github.com/aliereno/go-pagination/datatype"
	"github.com/aliereno/go-pagination/frameworks"
	"github.com/aliereno/go-pagination/pages"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   int
	Name string
}

func main() {
	app := fiber.New()

	db := initDB()
	items := initItems()

	app.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(pagination.Paginate(items, pagination.Config{
			PageType: pages.LinksPage{},
			Framework: frameworks.Fiber{
				Context: c,
			},
		}))
	})

	app.Get("/gorm", func(c *fiber.Ctx) error {
		query := db.Model(&User{}).Order("id desc")
		return c.JSON(pagination.Paginate(query, pagination.Config{
			Framework: frameworks.Fiber{
				Context: c,
			},
			Datatype: datatype.GORM{},
		}))
	})

	log.Fatal(app.Listen(":3000"))
}

func initItems() []interface{} {
	var items []interface{}
	for i := 1; i <= 101; i++ {
		if i%2 == 0 {
			items = append(items, i)
		} else {
			items = append(items, strconv.Itoa(i))
		}
	}
	return items
}
func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	var createUsers []User
	for i := 1; i <= 500; i++ {
		createUsers = append(createUsers, User{
			ID:   i,
			Name: "name " + strconv.Itoa(i),
		})
	}
	err = db.CreateInBatches(createUsers, 500).Error
	if err != nil {
		panic(err)
	}
	return db
}
