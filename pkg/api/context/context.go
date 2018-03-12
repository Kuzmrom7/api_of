package context

import (
	"context"
	"github.com/orderfood/api_of/pkg/api/config"
	_c "github.com/orderfood/api_of/pkg/common/context"
	"github.com/orderfood/api_of/pkg/storage"
)

var _ctx Context

type Context struct {
	_c.Context
	config  *config.Config
	storage storage.Storage
}

//-------------------------------------------------------------------------------------------------

func Get() *Context {
	return &_ctx
}

func (c *Context) SetConfig(cfg *config.Config) {
	c.config = cfg
}

func (c *Context) GetConfig() *config.Config {
	return c.config
}

func (c *Context) SetStorage(storage storage.Storage) {
	c.storage = storage
}

func (c *Context) GetStorage() storage.Storage {
	return c.storage
}

func (c *Context) Background() context.Context {
	return context.Background()
}
