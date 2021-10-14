package frameworks

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Fiber struct {
	Context *fiber.Ctx
}

func (fw Fiber) Check() error {
	if fw.Context == nil {
		return errors.New("fiber context cannot be nil")
	}
	return nil
}

func (fw Fiber) GetQueryParams() (page int, pageSize int) {
	p := new(params)
	if err := fw.Context.QueryParser(p); err != nil {
		return 1, 0
	}
	page, pageSize = p.Page, p.PageSize
	return
}

func (fw Fiber) GetURL() string {
	return fw.Context.BaseURL() + fw.Context.Path()
}
