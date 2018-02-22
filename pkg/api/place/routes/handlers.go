package routes

import (
	"net/http"
	"github.com/orderfood/api_of/pkg/log"

	"github.com/orderfood/api_of/pkg/api/place/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/util/http/utils"
)

//------------------------------------CREATE PLACE----------------------------------------------//
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

	plc, err := p.Create(usrid1, rq)
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

//------------------------------------LIST TYPE PLACE-------------------------------------------//
func TypePlaceList(w http.ResponseWriter, r *http.Request) {

	log.Debug("Handler: TypePlace: list type place")

	items, err := place.New(r.Context()).List()
	if err != nil {
		log.Errorf("Handler: TypePlace: list type place err ", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		log.Errorf("Handler: TypePlace: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: TypePlace: write response err: %s", err)
		return
	}
}

//------------------------------------INFORMATION ABOUT PLACE--------------------------------------------//
func GetPlaceInfo(w http.ResponseWriter, r *http.Request) {

	pid := utils.Vars(r)["place"]

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	log.Debug("Handler: Place: get place")

	p := place.New(r.Context())
	plc, err := p.GetPlaceByID(pid)
	if err != nil {
		log.Errorf("Handler: Place: get place", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if plc == nil {
		log.Warnf("Handler: Place: place by id `%s` not found", pid)
		errors.New("place").NotFound().Http(w)
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

//------------------------------------UPDATE PLACE--------------------------------------------//
func PlaceUpdate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	log.Debug("Handler: Place: update place")

	rq := new(request.RequestPlaceUpdate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	id := r.Context().Value("uid").(string)

	p := place.New(r.Context())

	plc, err := p.GetPlaceByID(rq.Id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if plc == nil {
		log.Warnf("Handler: Place: place by user id `%s` not found", id)
		errors.New("place").NotFound().Http(w)
		return
	}

	err = p.Update(plc, rq)
	if err != nil {
		log.Errorf("Handler: Place: update place err: %s", err)
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
