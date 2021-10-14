package pagination

import (
	"errors"

	"github.com/aliereno/go-pagination/datatype"
	"github.com/aliereno/go-pagination/frameworks"
	"github.com/aliereno/go-pagination/pages"
)

// Config defines the config for pagination.
type Config struct {
	// Default: 50
	PageSize int
	// Default: SimplePage
	PageType pages.IPage
	// Default: Array
	Datatype datatype.IDatabase
	// Default: Fiber
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
