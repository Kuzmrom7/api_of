package http

import (
	"net/http"
	"fmt"
	"github.com/orderfood/api_of/pkg/log"
)

func Handle(h http.HandlerFunc, middleware ...Middleware) http.HandlerFunc {
	headers := func(h http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			Headers(w, r)
			h.ServeHTTP(w, r)
		}
	}

	h = headers(h)
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func Listen(host string, port int, router http.Handler) error {

	log.Debugf("HTTP: listen HTTP server on %s:%d", host, port)

	return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router)
}
