package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/menu", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuCreate},
	{Path: "/menu", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetListMenu},
	{Path: "/menu/{menu}", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetInfoMenu},
	{Path: "/menu/{menu}/dish", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetListDishInMenu},
	{Path: "/menu/{menu}/ndish", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishListNotMenu},
	{Path: "/menu/{menu}/dish/{dish}", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishCreate},
	{Path: "/menu/{menu}/dish/{dish}", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishRemove},
}

//TODO удаление меню добавить
