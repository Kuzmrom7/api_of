package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	//{Path:"/dich", Method:http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler:GetList},
	{Path: "/dich", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DichCreate},
	{Path: "/dich", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DichRemove},
	{Path: "/dich", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DichList},
}
