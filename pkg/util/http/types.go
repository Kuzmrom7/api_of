package http

import "net/http"

const (
	MethodGet = http.MethodGet
	MethodPost = http.MethodPost
	MethodPut = http.MethodPut
	MethodDelete = http.MethodDelete

)

type Route struct {
	Path 		string
	Handler 	func(w http.ResponseWriter, r *http.Request)
	Middleware 	[]Middleware
	Method 		string
}

type Middleware func(http.HandlerFunc) http.HandlerFunc
