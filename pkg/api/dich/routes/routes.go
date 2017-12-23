package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	//{Path:"/dich", Method:http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler:GetList},
	{Path: "/dish", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishCreate},
	{Path: "/dish", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishRemove},
	//{Path: "/dish", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishGet},
	{Path: "/listdish", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishList},
	{Path: "/typedish", Method: http.MethodGet, Handler: TypeDishList},
}
