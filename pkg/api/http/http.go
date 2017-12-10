package http

import (
	"github.com/gorilla/mux"


	users "github.com/orderfood/api_of/pkg/auth/user/routes"
	"github.com/orderfood/api_of/pkg/util/http"
	session "github.com/orderfood/api_of/pkg/auth/session/routes"
	place "github.com/orderfood/api_of/pkg/api/place/routes"
	menu "github.com/orderfood/api_of/pkg/api/menu/routes"
	dich "github.com/orderfood/api_of/pkg/api/dich/routes"
)

var Routes = make([]http.Route, 0)

func AddRoutes(r ...[]http.Route) {
	for i := range r {
		Routes = append(Routes, r[i]...)
	}
}
func init() {

	AddRoutes(users.Routes)
	AddRoutes(session.Routes)
	AddRoutes(place.Routes)
	AddRoutes(menu.Routes)
	AddRoutes(dich.Routes)

}

func Listen(host string, port int){
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(http.Headers)

	for _, route := range Routes {
		r.Handle(route.Path, http.Handle(route.Handler, route.Middleware...)).Methods(route.Method)
	}

	http.Listen(host, port, r)
}