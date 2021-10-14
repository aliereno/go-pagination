package pagination

func Paginate(data interface{}, config ...Config) interface{} {
	// Set default config
	cfg := configDefault(config...)

	// Create Params and read it from query
	page, pageSize := cfg.Framework.GetQueryParams()

	// Check PageSize
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = cfg.PageSize
	}

	start := (page - 1) * pageSize
	end := page * pageSize

	items, totalSize, err := cfg.Datatype.GetSlicedItemsAndCount(data, start, end, pageSize)
	if err != nil {
		return err
	}
	totalPage := calculateTotalPage(totalSize, pageSize)

	return cfg.PageType.Response(page, pageSize, totalPage, items, cfg.Framework.GetURL())
}
