package pagination

import (
	"strconv"
	"testing"

	"github.com/aliereno/go-pagination/datatype"
	"github.com/aliereno/go-pagination/frameworks"
	"github.com/aliereno/go-pagination/pages"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/utils"
	"github.com/valyala/fasthttp"
)

// go test -run Test_Array_Paginate
func Test_Array_Paginate(t *testing.T) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)
	var testCases = []Config{
		{PageSize: 4, PageType: pages.SimplePage{}, Framework: frameworks.Fiber{Context: c}, Datatype: datatype.Array{}},
		{PageSize: 4, PageType: pages.SimplePage{}, Framework: frameworks.Fiber{Context: c}},
		{PageSize: 4, PageType: pages.LinksPage{}, Framework: frameworks.Fiber{Context: c}, Datatype: datatype.Array{}},
		{PageSize: 4, PageType: pages.LinksPage{}, Framework: frameworks.Fiber{Context: c}},
		{PageSize: 4, Framework: frameworks.Fiber{Context: c}, Datatype: datatype.Array{}},
		{PageSize: 4, Framework: frameworks.Fiber{Context: c}},
	}
	var items []interface{}
	for i := 1; i <= 5; i++ {
		if i%2 == 0 {
			items = append(items, i)
		} else {
			items = append(items, strconv.Itoa(i))
		}
	}
	for _, tt := range testCases {
		Paginate(items, tt)
	}
}

// go test -run=^$ -bench=Benchmark_String_Array -benchmem -count=4
func Benchmark_String_Array(b *testing.B) {
	app := fiber.New()
	c := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(c)

	var items []string
	for i := 1; i <= 10000; i++ {
		items = append(items, strconv.Itoa(i))
	}
	var err error
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		err = c.JSON(Paginate(items, Config{
			PageSize: 10,
		}))
	}
	utils.AssertEqual(b, nil, err)
	utils.AssertEqual(b, `{"currentPage":0,"pageSize":10,"totalPage":999,"items":["1","2","3","4","5","6","7","8","9","10"]}`, string(c.Response().Body()))
}
