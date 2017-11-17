package storage


import (
"github.com/orderfood/api_of/pkg/storage/pgsql"
"github.com/orderfood/api_of/pkg/storage/store"
)

func Get(c store.Config) (Storage, error) {
	var driver = ""
	switch driver {
	default:
		return pgsql.New(c)
	}
}
