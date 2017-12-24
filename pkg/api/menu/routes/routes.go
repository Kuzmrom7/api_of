package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	{Path: "/menu", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetListMenu},
	{Path: "/menu/info", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetMenu},
	{Path: "/menu", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuCreate},
	{Path: "/menudish", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishCreate},
	{Path: "/menudish", Method: http.MethodDelete, Middleware: []http.Middleware{middleware.Authenticate}, Handler: MenuDishRemove},
	{Path: "/menudish/list", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetFetchMenuDish},
	//TODO подумать как по другому назвать роуты, пока так
	{Path: "/menudish/listnotmenu", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: DishListNotMenu},

}

//TODO удаление меню добавить
