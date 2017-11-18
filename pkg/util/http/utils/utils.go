package utils

import (
	"context"
	"net/http"
)

func SetContext(r *http.Request, name string, val interface{}) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), name, val))
}
