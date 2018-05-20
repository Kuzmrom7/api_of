package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/adress", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: AdressCreate},
	{Path: "/adress/{place}", Method: http.MethodGet, Handler: GetListAdress},
}
