package routes

import (
	c "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/auth/session/routes/request"
	"github.com/orderfood/api_of/pkg/auth/session/views/v1"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
	"net/http"
)

func SessionCreate(w http.ResponseWriter, r *http.Request) {
	var (
		err     error
		storage = c.Get().GetStorage()
		usr     *types.User
	)

	rq := new(request.SessionCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Session: validation incoming data err: %v", err.Err())
		errors.New("incoming data invalid").Unknown().Http(w)
		return
	}

	usr, err = storage.User().GetByLogin(r.Context(), *rq.Login)
	if err == nil && usr == nil {
		log.Errorf("Handler: Session: account not found")
		errors.HTTP.Unauthorized(w)
		return
	}
	if err != nil {
		log.Errorf("Handler: Session: get account err: %v", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if err := usr.Security.Pass.ValidatePassword(*rq.Password); err != nil {
		log.Errorf("Handler: Session: check password err: %v", err)
		errors.HTTP.Unauthorized(w)
		return
	}

	session := types.NewSession(usr.Meta.ID, usr.Meta.Username, usr.Meta.Email)
	if err != nil {
		log.Errorf("Handler: Account: create account session failed")
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewSession(session).ToJson()
	if err != nil {
		log.Errorf("Handler: Session: convert struct to json err: %v", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Session: write response err: %v", err)
		return
	}
}
