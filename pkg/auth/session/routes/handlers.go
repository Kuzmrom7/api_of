package routes

import (
	c "github.com/orderfood/api_of/pkg/api/context"
	"net/http"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/auth/session/routes/request"
	"github.com/orderfood/api_of/pkg/auth/session/views/v1"
	"github.com/orderfood/api_of/pkg/common/errors"
	"log"
)

func SessionCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		storage = c.Get().GetStorage();
		usr     *types.User
	)

	rq := new(request.RequestSessionCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
	}

	usr, err = storage.User().GetByLogin(r.Context(), *rq.Login)
	if err == nil && usr == nil {
		errors.HTTP.Unauthorized(w)
		return
	}
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if err := usr.Security.Pass.ValidatePassword(*rq.Password); err != nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	session := types.NewSession(usr.Meta.ID, usr.Meta.Username, usr.Meta.Email)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewSession(session).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Session write response error")
		return
	}
}
