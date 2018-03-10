package routes

import (
	"net/http"

	"github.com/orderfood/api_of/pkg/api/menu/views/v1"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/api/menu/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/menu"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/api/dich"
	dv1 "github.com/orderfood/api_of/pkg/api/dich/views/v1"
	"github.com/orderfood/api_of/pkg/util/http/utils"
)

func MenuCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	log.Debug("Handler: Menu: create menu")

	rq := new(request.MenuCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Menu: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}

	men, err := menu.New(r.Context()).Create(rq)
	if err != nil {
		log.Errorf("Handler: Menu: create menu", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewMenu(men).ToJson()
	if err != nil {
		log.Errorf("Handler: Menu: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Menu: write response err: %s", err)
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

	log.Debug("Handler: Menu: List: get list menu")

	p := place.New(r.Context())
	pl, err := p.GetPlaceByIDUsr(id)
	if err != nil {
		log.Errorf("Handler: Menu: List: get place by user id %s err: %s", id, err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if pl == nil {
		log.Warnf("Handler: Menu: List: place by user id `%s` not found", id)
		errors.New("place").NotFound().Http(w)
		return
	}

	items, err := menu.New(r.Context()).List(pl.Meta.ID)
	if err != nil {
		log.Errorf("Handler: Menu: List: get list menu err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		log.Errorf("Handler: Menu: List: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Menu: List: write response err: %s", err)
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

	mid := utils.Vars(r)["menu"]
	did := utils.Vars(r)["dish"]

	log.Debug("Handler: Menu: Dish: add dish in menu")

	m := menu.New(r.Context())

	exists, err := m.CheckUniqueDishInMenu(mid, did)
	if err == nil && exists {
		log.Errorf("Handler: Menu: Dich: check unique, dish already adding")
		errors.New("menudish").NotUnique("dish").Http(w)
		return
	}
	if err != nil {
		log.Errorf("Handler: Menu: Dich: check exists by dish and menu err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	
	err = m.CreateMenuDish(mid, did)
	if err != nil {
		log.Errorf("Handler: Menu: Dish: add dish in menu err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte{}); err != nil {
		log.Errorf("Handler: Menu: Dish: write response err: %s", err)
		return
	}
}

func GetInfoMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	mid := utils.Vars(r)["menu"]

	var (
		err error
	)

	log.Debug("Handler: Menu: Info: get info menu")

	men, err := menu.New(r.Context()).GetMenuByID(mid)
	if err != nil {
		log.Errorf("Handler: Menu: Info: get info menu err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if men == nil {
		log.Warnf("Handler: Menu: Info: menu by id `%s` not found", mid)
		errors.New("menu").NotFound().Http(w)
		return
	}

	response, err := v1.NewMenu(men).ToJson()
	if err != nil {
		log.Errorf("Handler: Menu: info: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Menu: Info: write response err: %s", err)
		return
	}
}

func GetListDishInMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
		err error
	)

	mid := utils.Vars(r)["menu"]
	uid := r.Context().Value("uid").(string)

	log.Debug("Handler: Menu: Dish: List: get list dish in menu")

	dishList, err := menu.New(r.Context()).ListMenuDish(mid, uid)
	if err != nil {
		log.Errorf("Handler: Menu: Dish: List: get list dish in menu err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	typedishList, err := dich.New(r.Context()).TypeList()
	if err != nil {
		log.Errorf("Handler: Menu: Dish: List: get type dishes err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := dv1.NewListD(dishList, typedishList).ToJson()
	if err != nil {
		log.Errorf("Handler: Menu: Dish: List: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Menu: Dish: List: write response err: %s", err)
		return
	}

}

func MenuDishRemove(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	mid := utils.Vars(r)["menu"]
	did := utils.Vars(r)["dish"]

	log.Debug("Handler: Menu: Dish: delete dish from menu")

	err := menu.New(r.Context()).RemoveMenuDish(mid, did)
	if err != nil {
		log.Errorf("Handler: Menu: Dish: delete dish from menu err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte{}); err != nil {
		log.Errorf("Handler: Menu: Dish: write response err: %s", err)
		return
	}
}

func DishListNotMenu(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	usrid := r.Context().Value("uid").(string)
	mid := utils.Vars(r)["menu"]

	log.Debug("Handler: Menu: Dish: List: get list dish not menu")

	dishList, err := menu.New(r.Context()).ListDishNotMenu(mid, usrid)
	if err != nil {
		log.Errorf("Handler: Menu: Dish: List: get list dish not menu err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := dv1.NewList(dishList).ToJson()
	if err != nil {
		log.Errorf("Handler: Menu: Dish: List: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Menu: Dish: List: write response err: %s", err)
		return
	}
}
