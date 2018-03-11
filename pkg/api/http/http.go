package http

import (
	"github.com/gorilla/mux"

	adress "github.com/orderfood/api_of/pkg/api/adress/routes"
	dich "github.com/orderfood/api_of/pkg/api/dich/routes"
	menu "github.com/orderfood/api_of/pkg/api/menu/routes"
	personal "github.com/orderfood/api_of/pkg/api/personal/routes"
	place "github.com/orderfood/api_of/pkg/api/place/routes"
	session "github.com/orderfood/api_of/pkg/auth/session/routes"
	users "github.com/orderfood/api_of/pkg/auth/user/routes"
	"github.com/orderfood/api_of/pkg/util/http"
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
	AddRoutes(personal.Routes)
	AddRoutes(adress.Routes)

}

func Listen(host string, port int) {
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(http.Headers)

	for _, route := range Routes {
		r.Handle(route.Path, http.Handle(route.Handler, route.Middleware...)).Methods(route.Method)
	}

	http.Listen(host, port, r)
}
