package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path: "/place", Method: http.MethodPost, Middleware: []http.Middleware{middleware.Authenticate}, Handler: PlaceCreate},
	{Path: "/place", Method: http.MethodPut, Middleware: []http.Middleware{middleware.Authenticate}, Handler: PlaceUpdate},
	{Path: "/place", Method: http.MethodGet, Middleware: []http.Middleware{middleware.Authenticate}, Handler: GetPlaceInfo},
	{Path: "/type/place", Method: http.MethodGet, Handler: TypePlaceList},
}
