package routes

import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/api/dich/views/v1"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/dich"
)

func DishCreate(w http.ResponseWriter, r *http.Request) {

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

	typedish_id, err := d.GetIDTypeDishByName(rq.TypeDish)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(typedish_id)
	if typedish_id == "" {
		errors.New("typedish").NotFound().Http(w)
	}

	usrid1 := r.Context().Value("uid").(string)

	di, err := d.Create(rq, typedish_id, usrid1)
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

func DishUpdate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestDichUpdate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid1 := r.Context().Value("uid").(string)

	dish, err := dich.New(r.Context()).GetDishByUsrIdAndDishName(usrid1, *rq.Name)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if dish == nil {
		errors.New("dish").NotFound().Http(w)
	}

	err = dich.New(r.Context()).Update(usrid1, rq, dish)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewDich(dish).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Dish update response error")
		return
	}

}

func DishGet(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	rq := new(request.RequestDichRemoveFetch)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	dish, err := dich.New(r.Context()).GetDishByUsrIdAndDishName(usrid, rq.Name)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if dish == nil {
		errors.New("dish").NotFound().Http(w)
	}

	response, err := v1.NewDich(dish).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Dish fetch response error")
		return
	}
}

func DishRemove(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	rq := new(request.RequestDichRemoveFetch)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	d := dich.New(r.Context())

	dich_id, err := d.GetIDdishByName(rq.Name, usrid)
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

func DishList(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	usrid1 := r.Context().Value("uid").(string)

	items, err := dich.New(r.Context()).List(usrid1)
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

func TypeDishList(w http.ResponseWriter, r *http.Request) {

	items, err := dich.New(r.Context()).TypeList()
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
		log.Println("Dich list response error")
		return
	}
}
