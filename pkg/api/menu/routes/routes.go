package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	{Path: "/menu", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetListMenu},
	{Path: "/menu", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuCreate},
}

//TODO удаление добавить
