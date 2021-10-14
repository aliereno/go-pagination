package frameworks

type IFramework interface {
	Check() error
	GetQueryParams() (int, int)
	GetURL() string
}

// Params stands for query params.
type params struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}
