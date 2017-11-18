package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)
var Routes = []http.Route{

	{Path:"/user", Method:http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler:GetUser},
	{Path:"/user", Method:http.MethodPost, Handler:UserCreate},

}