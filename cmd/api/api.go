package main

import (
	"github.com/gorilla/mux"

	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/storage"

	users "github.com/orderfood/api_of/pkg/api/users/routes"
)

const (
	host = "localhost"
	port = 8080
)

var Routes = make([]http.Route, 0)

func AddRoutes(r ...[]http.Route) {
	for i := range r {
		Routes = append(Routes, r[i]...)
	}
}
func init() {

	AddRoutes(users.Routes)
}

func main() {

	storage.DB_connect()

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(http.Headers)

	for _, route := range Routes {
		r.Handle(route.Path, http.Handle(route.Handler, route.Middleware...)).Methods(route.Method)
	}

	http.Listen(host, port, r)
}
