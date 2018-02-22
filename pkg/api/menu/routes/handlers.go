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

	rq := new(request.MenuCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	pl, err := place.New(r.Context()).GetPlaceByIDUsr(usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	if pl == nil {
		errors.New("place").NotFound().Http(w)
	}

	men, err := menu.New(r.Context()).Create(pl.Meta.ID, rq)
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	response, err := v1.NewMenu(men).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

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
	pl, err := p.GetPlaceByIDUsr(id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if pl == nil {
		errors.New("place").NotFound().Http(w)
	}

	items, err := menu.New(r.Context()).List(pl.Meta.ID)
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

	rq := new(request.RequestMenuDishCreateRemove)
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

	w.WriteHeader(http.StatusOK)

}

func GetMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	rq := new(request.RequestMenuFetch)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	pl, err := place.New(r.Context()).GetPlaceByIDUsr(usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	if pl == nil {
		errors.New("place").NotFound().Http(w)
	}

	men, err := menu.New(r.Context()).GetMenuByIDPlaceAndNameMenu(pl.Meta.ID, rq.NameMenu)
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

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Menu write response error")
		return
	}
}

func GetFetchMenuDish(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	rq := new(request.RequestMenuFetch)
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

	usrid := r.Context().Value("uid").(string)

	dishList, err := menu.New(r.Context()).ListMenuDish(menu_id, usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	typedishList, err := dich.New(r.Context()).TypeList()
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := dv1.NewListD(dishList, typedishList).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("MenuDish list response error")
		return
	}

}

func MenuDishRemove(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	usrid := r.Context().Value("uid").(string)

	rq := new(request.RequestMenuDishCreateRemove)
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

	err = m.RemoveMenuDish(menu_id, dish_id)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte{}); err != nil {
		log.Println("MenuDich remove response error")
		return
	}
}

func DishListNotMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	rq := new(request.RequestMenuFetch)
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

	usrid := r.Context().Value("uid").(string)

	dishList, err := menu.New(r.Context()).ListDishNotMenu(menu_id, usrid)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := dv1.NewList(dishList).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Dich list response error")
		return
	}
}
