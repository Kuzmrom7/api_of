package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/menu", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuCreate},
	{Path: "/menu/{menu}", Method: http.MethodGet, Handler: GetInfoMenu},
	{Path: "/menu/place/{place}", Method: http.MethodGet, Handler: GetListMenu},
	{Path: "/menu/{menu}/place/{place}/dish", Method: http.MethodGet, Handler: GetListDishInMenu},
	{Path: "/menu/{menu}/place/{place}/ndish", Method: http.MethodGet, Handler: DishListNotMenu},
	{Path: "/menu/{menu}/dish/{dish}", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishCreate},
	{Path: "/menu/{menu}/dish/{dish}", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishRemove},
}

//TODO удаление меню добавить
