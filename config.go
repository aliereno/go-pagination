package pagination

import (
	"errors"

	"github.com/aliereno/go-pagination/datatype"
	"github.com/aliereno/go-pagination/frameworks"
	"github.com/aliereno/go-pagination/pages"
)

// Config defines the config for pagination.
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

// ConfigDefault is the default config
var ConfigDefault = Config{
	PageSize: 50,
	PageType: pages.SimplePage{},
	Datatype: datatype.Array{},
}

// Helper function to set default values
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		panic(errors.New("framework context can not be null"))
		//return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default values

	if cfg.PageSize <= 0 {
		cfg.PageSize = ConfigDefault.PageSize
	}
	if cfg.PageType == nil {
		cfg.PageType = ConfigDefault.PageType
	}
	if cfg.Datatype == nil {
		cfg.Datatype = ConfigDefault.Datatype
	}
	if cfg.Framework == nil {
		panic(errors.New("framework cannot be null"))
	}
	if err := cfg.Framework.Check(); err != nil {
		panic(err)
	}
	return cfg
}
