package config

import "github.com/orderfood/api_of/pkg/storage/store"

var _cfg = new(Config)

func Set(cfg *Config) *Config {
	_cfg = cfg
	return _cfg
}

func Get() *Config {
	return _cfg
}

func (c *Config) GetPGDB() store.Config {
	return store.Config{
		Connection: *c.Database.Connection,
	}
}