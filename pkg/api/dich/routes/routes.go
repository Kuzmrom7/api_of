package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/dish", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishCreate},
	{Path: "/dish", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishUpdate},
	{Path: "/dish/place/{place}", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishList},
	{Path: "/dish/{dish}", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishRemove},
	{Path: "/dish/{dish}", Method: http.MethodGet, Handler: DishGet},
	{Path: "/type/dish", Method: http.MethodGet, Handler: TypeDishList},
}

//TODO обновление и fethc добавить
