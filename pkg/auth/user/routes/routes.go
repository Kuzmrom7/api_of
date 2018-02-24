package routes

import (
	"github.com/orderfood/api_of/pkg/util/http"
)

var Routes = []http.Route{
	{Path: "/user", Method: http.MethodPost, Handler: UserCreate},
}
