package routes

import (
	"net/http"

	"github.com/orderfood/api_of/pkg/api/adress"
	"github.com/orderfood/api_of/pkg/api/adress/routes/request"
	"github.com/orderfood/api_of/pkg/api/adress/views/v1"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/util/http/utils"
)

func AdressCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	log.Debug("Handler: Adress: create adress")

	rq := new(request.AdressCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Adress: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}

	adr, err := adress.New(r.Context()).Create(rq)
	if err != nil {
		log.Errorf("Handler: Adress: create adress", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewAdress(adr).ToJson()
	if err != nil {
		log.Errorf("Handler: Adress: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Adress: write response err: %s", err)
		return
	}
}

func GetListAdress(w http.ResponseWriter, r *http.Request) {

	var (
		err error
		pid = utils.Vars(r)["place"]
	)

	log.Debug("Handler: Adress: List: get list adress")

	p := place.New(r.Context())
	pl, err := p.GetPlaceByID(pid)
	if err != nil {
		log.Errorf("Handler: Adress: List: get place by user id %s err: %s", pid, err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if pl == nil {
		log.Warnf("Handler: Adress: List: place by user id `%s` not found", pid)
		errors.New("place").NotFound().Http(w)
		return
	}

	items, err := adress.New(r.Context()).List(pl.Meta.ID)
	if err != nil {
		log.Errorf("Handler: Adress: List: get list adress err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		log.Errorf("Handler: Adress: List: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Adress: List: write response err: %s", err)
		return
	}

}
