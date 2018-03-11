package utils

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
)

func SetContext(r *http.Request, name string, val interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), name, val))
}

func Vars(r *http.Request) map[string]string {
	return mux.Vars(r)
}
