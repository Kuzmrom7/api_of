package routes


import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	//{Path:"/place", Method:http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler:GetPlace},
	{Path: "/personal", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: PersonCreate},
	//{Path: "/typepersonal", Method: http.MethodGet, Handler: TypePersonList},
}

