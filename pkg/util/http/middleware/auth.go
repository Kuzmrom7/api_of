package middleware

import (
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/util/http/utils"
	"net/http"
	"strings"
)

func Authenticate(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var token string

		if r.Header.Get("Authorization") != "" {
			//parse Authorization header
			var auth = strings.SplitN(r.Header.Get("Authorization"), " ", 2)

			if len(auth) != 2 || auth[0] != "Bearer" {
				errors.HTTP.Unauthorized(w)
				return
			}
			token = auth[1]
		} else {
			errors.HTTP.Unauthorized(w)
			return
		}

		s := new(types.Session)
		err := s.Decode(token)
		if err != nil {
			errors.HTTP.Unauthorized(w)
			return
		}

		// set user data in request context
		r = utils.SetContext(r, "uid", s.Uid)
		r = utils.SetContext(r, "username", s.Username)
		r = utils.SetContext(r, "email", s.Email)

		h.ServeHTTP(w, r)
	}
}
