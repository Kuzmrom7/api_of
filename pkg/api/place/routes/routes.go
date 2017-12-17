package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{

	//{Path:"/place", Method:http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler:GetPlace},
	{Path:"/place", Method:http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler:PlaceCreate},
	{Path: "/place", Method: http.MethodGet, Handler: PlaceList},
}