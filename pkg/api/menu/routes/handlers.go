package routes

import (
	"net/http"

	"github.com/orderfood/api_of/pkg/api/menu/views/v1"
	"github.com/orderfood/api_of/pkg/api/menu/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/menu"
	"log"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/api/dich"
	"github.com/orderfood/api_of/pkg/api/personal"
	dv1 "github.com/orderfood/api_of/pkg/api/dich/views/v1"
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

	place_id, err := m.GetIDplaceByName(rq.NamePlace)
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

func GetListMenu(w http.ResponseWriter, r *http.Request) {

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

	items, err := menu.New(r.Context()).List(plc.Meta.ID)
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

func MenuDishCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	rq := new(request.RequestMenuDishCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	m := menu.New(r.Context())

	menu_id, err := m.GetIDMenuByName(rq.NameMenu)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(menu_id)
	if menu_id == "" {
		errors.New("menu").NotFound().Http(w)
	}

	dish_id, err := dich.New(r.Context()).GetIDdishByName(rq.NameDish, usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(dish_id)
	if dish_id == "" {
		errors.New("dish").NotFound().Http(w)
	}

	err = m.CreateMenuDish(menu_id, dish_id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}
	//
	//response, err := v1.NewMenu(men).ToJson()
	//if err != nil {
	//	errors.HTTP.InternalServerError(w)
	//}

	w.WriteHeader(http.StatusOK)
	//if _, err = w.Write(response); err != nil {
	//	log.Println("Menu write response error")
	//	return
	//}
}

func GetMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
		id  = r.Context().Value("uid").(string)
	)

	rq := new(request.RequestMenuFetch)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	p := personal.New(r.Context())
	place_id, err := p.GetIDPlaceByUsr(id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(place_id)
	if place_id == "" {
		errors.New("place").NotFound().Http(w)
	}

	m := menu.New(r.Context())
	men, err := m.GetMenuByIDPlaceAndNameMenu(place_id, rq.Name)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if men == nil {
		errors.New("menu").NotFound().Http(w)
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

func GetListMenuDish(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	rq := new(request.RequestMenuDishList)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	m := menu.New(r.Context())

	menu_id, err := m.GetIDMenuByName(rq.NameMenu)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(menu_id)
	if menu_id == "" {
		errors.New("menu").NotFound().Http(w)
	}

	d := dich.New(r.Context())

	typedish_id, err := d.GetIDTypeDishByName(rq.NameTypeDish)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	log.Print(typedish_id)
	if typedish_id == "" {
		errors.New("typedish").NotFound().Http(w)
	}

	items, err := menu.New(r.Context()).ListMenuDish(menu_id, typedish_id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := dv1.NewList(items).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("MenuDish list response error")
		return
	}

}