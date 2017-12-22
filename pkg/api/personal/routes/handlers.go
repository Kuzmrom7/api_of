package routes
import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/api/personal/views/v1"
	"github.com/orderfood/api_of/pkg/api/personal/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/personal"

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

	p := personal.New(r.Context())

	typeperson_id, err := p.GetIDTypePersonByName(rq.NameTypePerson)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if typeperson_id == "" {
		errors.New("type_personal").NotFound().Http(w)
	}

	usrid1 := r.Context().Value("uid").(string)

	place_id, err := p.GetIDPlaceByUsr(usrid1)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(place_id)
	if place_id == "" {
		errors.New("place").NotFound().Http(w)
	}


	pers, err := p.Create(typeperson_id, place_id, rq)
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

	items, err := personal.New(r.Context()).List()
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
		log.Println("Personal list response error")
		return
	}
}
