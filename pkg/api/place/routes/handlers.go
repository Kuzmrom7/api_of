package routes

import (
	"net/http"
	"github.com/orderfood/api_of/pkg/log"

	"github.com/orderfood/api_of/pkg/api/place/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/place"
	"fmt"
)

//------------------------------------СОЗДАНИЕ ЗАВЕДЕНИЯ----------------------------------------------//
func PlaceCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	log.Debug("Handler: Place: create place")

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestPlaceCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Place: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}
	usrid1 := r.Context().Value("uid").(string)

	p := place.New(r.Context())

	typeplace_id, err := p.GetIDTypePlaceByName(rq.NameTypePlace)
	if err != nil {
		log.Errorf("Handler: Place: get id type place by type name err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if typeplace_id == "" {
		log.Warnf("Handler: Place: id type place by name `%s` not found", rq.NameTypePlace)
		errors.New("type_place").NotFound().Http(w)
		return
	}

	plc, err := p.Create(usrid1, typeplace_id, rq)
	if err != nil {
		log.Errorf("Handler: Place: create place", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		log.Errorf("Handler: Place: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Place: write response err: %s", err)
		return
	}
}

//------------------------------------СПИСОК ТИПОВ ЗАВЕДЕНИЙ--------------------------------------------//
func TypePlaceList(w http.ResponseWriter, r *http.Request) {

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
func PlaceUpdate(w http.ResponseWriter, r *http.Request) {

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
	fmt.Println("-----")
	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}
	fmt.Println("---+++++++++++++--")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Place write response error")
		return
	}

}
