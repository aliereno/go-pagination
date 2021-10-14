package main

import (
	"log"
	"strconv"

	"github.com/aliereno/go-pagination"
	"github.com/aliereno/go-pagination/frameworks"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	var items []interface{}
	for i := 1; i <= 101; i++ {
		if i%2 == 0 {
			items = append(items, i)
		} else {
			items = append(items, strconv.Itoa(i))
		}
	}

	app.Get("/items", func(c *fiber.Ctx) error {
		return c.JSON(pagination.Paginate(items, pagination.Config{
			Framework: frameworks.Fiber{
				Context: c,
			},
		}))
	})

	log.Fatal(app.Listen(":3000"))
}
