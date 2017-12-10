package routes

import (
	"net/http"

	"github.com/orderfood/api_of/pkg/api/menu/views/v1"
	"github.com/orderfood/api_of/pkg/api/menu/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/menu"
	"log"
)

func MenuCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestMenuCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	m := menu.New(r.Context())

	place_id, err := m.GetIDByName(rq.NamePlace)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(place_id)
	if place_id == "" {
		errors.New("place").NotFound().Http(w)
	}

	men, err := m.Create(place_id, rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewMenu(men).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	///log.Println("Create user id: " , usr.Meta.ID, " username: " , usr.Meta.Username)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Menu write response error")
		return
	}
}
