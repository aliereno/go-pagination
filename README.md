# Go Pagination
<p align="center">
    <a href="https://pkg.go.dev/github.com/aliereno/go-pagination" target="_blank"><img src="https://img.shields.io/github/go-mod/go-version/aliereno/go-pagination?style=for-the-badge&logo=go" alt="golang version" /></a>&nbsp;
    <a href="https://goreportcard.com/report/github.com/aliereno/go-pagination" target="_blank"><img src="https://goreportcard.com/badge/github.com/aliereno/go-pagination?style=for-the-badge&logo=none" alt="go report" /></a>&nbsp;
    <a href="https://github.com/aliereno/go-pagination/blob/master/LICENSE" target="_blank"><img src="https://img.shields.io/badge/LICENSE-MIT-red?style=for-the-badge&logo=none" alt="go-pagination license" /></a>
</p>

*** This repo is written with learning and experimenting purposes. I am open to suggestions or any kind of help. ***

## Installation

```bash
go get -u github.com/aliereno/go-pagination
```


Database integrations:

- [X] [gorm](https://github.com/go-gorm/gorm)

Framework integrations:

- [X] [fiber](https://github.com/gofiber/fiber)

## Example

```go
func main() {
	app := fiber.New()

	// paginate simple array
	app.Get("/array", func(c *fiber.Ctx) error {
		return c.JSON(pagination.Paginate(items, pagination.Config{
			Framework: frameworks.Fiber{
				Context: c,
			},
		}))
	})

	// paginate gorm.DB query
	app.Get("/gorm", func(c *fiber.Ctx) error {
		query := gorm.DB.Model(&User{}).Order("id desc")
		return c.JSON(pagination.Paginate(query, pagination.Config{
			PageSize: 100,
			PageType: pages.LinksPage{},
			Datatype: datatype.GORM{},
			Framework: frameworks.Fiber{
				Context: c,
			},
		}))
	})

	log.Fatal(app.Listen(":3000"))
}
```

## Config

```go
type Config struct {
	
	// Max number of items per page
	//
	// Default: 50
	PageSize int

	// Structure of pagination response
	//
	// Default: SimplePage | Choices: SimplePage, LinksPage
	PageType pages.IPage
	
	// Type of variable to be given to the Paginate func
	//
	// Default: Array | Choices: Array, GORM
	Datatype datatype.IDatatype
	
	// Framework to pull page and page_size variables
	//
	// Required | Choices: Fiber
	Framework frameworks.IFramework
}
```

## Customization
You can customize the response structure.
```go
type CustomPage struct {
	CustomField string      `json:"customField"`
	Page        int         `json:"currentPage"`
	Items       interface{} `json:"items"`
}

func (s CustomPage) Response(page int, pageSize int, totalPage int, items interface{}, query string) interface{} {
	s.Page = page
	s.Items = items
	return s
}
```
You can use it in the Paginate function.
```go
pagination.Paginate(items, pagination.Config{
	PageType: CustomPage{
		CustomField: time.Now().String(),
	},
	Framework: frameworks.Fiber{
		Context: c,
	},
})
```