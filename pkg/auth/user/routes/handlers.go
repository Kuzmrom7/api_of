package routes

import (
	"net/http"
	"log"

	sv1 "github.com/orderfood/api_of/pkg/auth/session/views/v1"
	"github.com/orderfood/api_of/pkg/auth/user/views/v1"
	"github.com/orderfood/api_of/pkg/auth/user/routes/request"
	"github.com/orderfood/api_of/pkg/auth/user"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/common/errors"

)

func GetUser (w http.ResponseWriter, r *http.Request){

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err      error
		id = r.Context().Value("uid").(string)
	)

	u := user.New(r.Context())
	usr, err := u.GetByID(id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if usr == nil {
		errors.New("user").NotFound().Http(w)
	}

	response, err := v1.NewUser(usr).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		return
	}

}

func UserCreate (w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	rq := new(request.RequestUserCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	u := user.New(r.Context())

	exists, err := u.CheckExists(*rq.Username)
	if err == nil && exists {
		errors.New("user").NotUnique("username").Http(w)
		return
	}
	if err != nil{
		errors.HTTP.InternalServerError(w)
	}
	exists, err = u.CheckExists(*rq.Email)
	if err == nil && exists {
		errors.New("user").NotUnique("email").Http(w)
		return
	}
	if err != nil{
		errors.HTTP.InternalServerError(w)
	}

	usr, err := u.Create(rq)

	session := types.NewSession(usr.Meta.ID, usr.Meta.Username, usr.Meta.Email)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := sv1.NewSession(session).ToJson()
	if err != nil{
		errors.HTTP.InternalServerError(w)
		return
	}

	log.Println("Create user id: " , usr.Meta.ID, " username: " , usr.Meta.Username)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil{
		log.Println("User write response error")
		return
	}
}