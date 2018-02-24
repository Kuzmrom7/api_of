package routes

import (
	"net/http"

	sv1 "github.com/orderfood/api_of/pkg/auth/session/views/v1"
	"github.com/orderfood/api_of/pkg/auth/user/routes/request"
	"github.com/orderfood/api_of/pkg/auth/user"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	log.Debug("Handler: User: create user")

	rq := new(request.RequestUserCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: User: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}

	u := user.New(r.Context())

	exists, err := u.CheckExists(*rq.Username)
	if err == nil && exists {
		log.Errorf("Handler: User: username `%s` not unique", err)
		errors.New("user").NotUnique("username").Http(w)
		return
	}
	if err != nil {
		log.Errorf("Handler: User: check exists by username err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	exists, err = u.CheckExists(*rq.Email)
	if err == nil && exists {
		log.Errorf("Handler: User: email `%s` not unique", err)
		errors.New("user").NotUnique("email").Http(w)
		return
	}
	if err != nil {
		log.Errorf("Handler: User: check exists by email err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	usr, err := u.Create(rq)

	session := types.NewSession(usr.Meta.ID, usr.Meta.Username, usr.Meta.Email)
	if err != nil {
		log.Errorf("Handler: User: create user", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := sv1.NewSession(session).ToJson()
	if err != nil {
		log.Errorf("Handler: User: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: User: write response err: %s", err)
		return
	}
}
