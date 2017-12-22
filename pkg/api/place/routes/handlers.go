package routes

import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/api/place/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/place"
)

//------------------------------------СОЗДАНИЕ ЗАВЕДЕНИЯ----------------------------------------------//
func PlaceCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestPlaceCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}
	usrid1 := r.Context().Value("uid").(string)

	p := place.New(r.Context())

	typeplace_id, err := p.GetIDTypePlaceByName(rq.NameTypePlace)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if typeplace_id == "" {
		errors.New("type_place").NotFound().Http(w)
	}

	plc, err := p.Create(usrid1, typeplace_id, rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	///log.Println("Create user id: " , usr.Meta.ID, " username: " , usr.Meta.Username)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Place write response error")
		return
	}
}

//------------------------------------СПИСОК ТИПОВ ЗАВЕДЕНИЙ--------------------------------------------//
func PlaceList(w http.ResponseWriter, r *http.Request) {

	items, err := place.New(r.Context()).List()
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
		log.Println("Dich list response error")
		return
	}
}

//------------------------------------ИНФОРМАЦИЯ О ЗАВЕДЕНИИ--------------------------------------------//
func GetPlace(w http.ResponseWriter, r *http.Request) {

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

	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		return
	}
}
//------------------------------------ОБНОВЛЕНИЕ ЗАВЕДЕНИЯ--------------------------------------------//
func PlaceUpdate(w http.ResponseWriter, r *http.Request){

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestPlaceUpdate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid1 := r.Context().Value("uid").(string)

	p := place.New(r.Context())

	plc, err := p.GetPlaceByIDUsr(usrid1)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if plc == nil {
		errors.New("place").NotFound().Http(w)
	}

	err = p.Update(plc, rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Place write response error")
		return
	}

}
