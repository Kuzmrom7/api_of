package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
	"github.com/orderfood/api_of/pkg/util/http/middleware"
)

var Routes = []http.Route{
	{Path:"/place", Method:http.MethodGet,[]http.Middleware{middleware.Authenticate}, Handler: GetPlace},

}

