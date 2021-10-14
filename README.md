# Go Pagination

*** This repo is written with learning and experimenting purposes. I am open to suggestions or any kind of help. ***

## Installation

```bash
go get -u github.com/aliereno/go-pagination
```


Database integrations:

- [ ] [gorm](https://github.com/go-gorm/gorm)

Framework integrations:

- [X] [fiber](https://github.com/gofiber/fiber)

## Example

```go
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
	for i := 1; i <= 15; i++ {
        items = append(items, i)
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
```
