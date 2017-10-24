package routes

import "github.com/orderfood/api_of/pkg/util/http"

var Routes = []http.Route{
	{Path:"/users", Method:http.MethodGet, Handler:GetUser},
}