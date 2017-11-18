package routes

import "github.com/orderfood/api_of/pkg/util/http"

var Routes = []http.Route{
	//session handlers
	{Path: "/session", Method: http.MethodPost, Handler: SessionCreate},
}
