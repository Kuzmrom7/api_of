package main

import (
	"github.com/orderfood/api_of/pkg/storage"
	"github.com/orderfood/api_of/pkg/api"
)

const (
	host = "localhost"
	port = 8080
)


func main() {

	storage.DB_connect()

	api.Daemon(host,port)


}
