package routes

import (
	"log"
	"net/http"

	"github.com/orderfood/api_of/pkg/api/personal"
	"github.com/orderfood/api_of/pkg/api/personal/routes/request"
	"github.com/orderfood/api_of/pkg/api/personal/views/v1"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/common/errors"
)

func PersonCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestPersonCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	place, err := place.New(r.Context()).GetPlaceByIDUsr(usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	if place == nil {
		errors.New("place").NotFound().Http(w)
	}

	p := personal.New(r.Context())

	typeperson_id, err := p.GetIDTypePersonByName(rq.NameTypePerson)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if typeperson_id == "" {
		errors.New("type_personal").NotFound().Http(w)
	}

	pers, err := p.Create(typeperson_id, place.Meta.ID, rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewPersonal(pers).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Personal write response error")
		return
	}
}

func TypePersonList(w http.ResponseWriter, r *http.Request) {

	items, err := personal.New(r.Context()).ListType()
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewListType(items).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Personal list response error")
		return
	}
}

func GetListPersonal(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
		id  = r.Context().Value("uid").(string)
	)

	p := place.New(r.Context())
	plc, err := p.GetPlaceByIDUsr(id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if plc == nil {
		errors.New("place").NotFound().Http(w)
	}

	items, err := personal.New(r.Context()).List(plc.Meta.ID)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Menu list response error")
		return
	}

}
