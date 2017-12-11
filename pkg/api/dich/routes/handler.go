package routes

import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/api/dich/views/v1"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/dich"
	//"github.com/orderfood/api_of/pkg/util/http/utils"

)

func DichCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestDichCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	d := dich.New(r.Context())

	di, err := d.Create(rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewDich(di).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Dich write response error")
		return
	}
}

func DichRemove(w http.ResponseWriter, r *http.Request) {

	//nameDich := utils.Vars(r)["dich"]

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestDichRemove)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	d := dich.New(r.Context())

	dich_id, err := d.GetIDByName(rq.Name)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(dich_id)
	if dich_id == "" {
		errors.New("dich").NotFound().Http(w)
	}

	err = d.Remove(dich_id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte{}); err != nil {
		log.Println("Dich remove response error")
		return
	}
}

func DichList(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	items, err := dich.New(r.Context()).List()
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
