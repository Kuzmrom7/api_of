package routes

import (
	"net/http"

	"github.com/orderfood/api_of/pkg/api/dich"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/api/dich/views/v1"
	"github.com/orderfood/api_of/pkg/api/place"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/util/http/utils"
)

func DishCreate(w http.ResponseWriter, r *http.Request) {

	var (
		err    error
		usrid1 = r.Context().Value("uid").(string)
	)

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	log.Debug("Handler: Dish: create dish")

	rq := new(request.DishCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Dish: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}

	p := place.New(r.Context())
	plc, err := p.GetPlaceByIDUsr(usrid1)
	if err != nil {
		log.Errorf("Handler: Place: get place", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if plc == nil {
		log.Warnf("Handler: Place: place by id user `%s` not found", usrid1)
		errors.New("place").NotFound().Http(w)
		return
	}

	di, err := dich.New(r.Context()).Create(rq, plc.Meta.ID)
	if err != nil {
		log.Errorf("Handler: Dish: create dish", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewDich(di).ToJson()
	if err != nil {
		log.Errorf("Handler: Dish: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Dish: write response err: %s", err)
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

	log.Debug("Handler: Dish: update dish")

	rq := new(request.DishUpdate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		log.Errorf("Handler: Dish: validation incoming data err: %s", err.Err())
		errors.New("Invalid incoming data").Unknown().Http(w)
		return
	}

	usrid1 := r.Context().Value("uid").(string)

	dish, err := dich.New(r.Context()).GetDishById(rq.Id)
	if err != nil {
		log.Errorf("Handler: Dish: get dish by id %s err: %s", rq.Id, err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if dish == nil {
		log.Warnf("Handler: Dish: dish by id `%s` not found", rq.Id)
		errors.New("place").NotFound().Http(w)
		return
	}

	err = dich.New(r.Context()).Update(usrid1, rq, dish)
	if err != nil {
		log.Errorf("Handler: Dish: update dish err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewDich(dish).ToJson()
	if err != nil {
		log.Errorf("Handler: Dish: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Dish: write response err: %s", err)
		return
	}

}

func DishGet(w http.ResponseWriter, r *http.Request) {

	var (
		err error
		did = utils.Vars(r)["dish"]
	)

	log.Debug("Handler: Dish: get dish")

	dish, err := dich.New(r.Context()).GetDishById(did)
	if err != nil {
		log.Errorf("Handler: Dish: get dish", err)
		errors.HTTP.InternalServerError(w)
		return
	}
	if dish == nil {
		log.Warnf("Handler: Dish: dish by id `%s` not found", did)
		errors.New("dish").NotFound().Http(w)
		return
	}

	response, err := v1.NewDich(dish).ToJson()
	if err != nil {
		log.Errorf("Handler: Dish: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Dish: write response err: %s", err)
		return
	}
}

func DishRemove(w http.ResponseWriter, r *http.Request) {

	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	did := utils.Vars(r)["dish"]

	log.Debug("Handler: Dish: delete dish")

	err := dich.New(r.Context()).Remove(did)
	if err != nil {
		log.Errorf("Handler: Dish: delete dish err %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte{}); err != nil {
		log.Errorf("Handler: Dish: Dish: write response err: %s", err)
		return
	}
}

func DishList(w http.ResponseWriter, r *http.Request) {

	log.Debug("Handler: Dish: List: list dishes")
	var (
		err error
		pid = utils.Vars(r)["place"]
	)

	items, err := dich.New(r.Context()).List(pid)
	if err != nil {
		log.Errorf("Handler: Dish: List: list dishes err ", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		log.Errorf("Handler: Dish: List: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: Dish: List: write response err: %s", err)
		return
	}
}

func TypeDishList(w http.ResponseWriter, r *http.Request) {

	log.Debug("Handler: TypeDish: list type dish")

	items, err := dich.New(r.Context()).TypeList()
	if err != nil {
		log.Errorf("Handler: TypeDish: list type dish err ", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewListType(items).ToJson()
	if err != nil {
		log.Errorf("Handler: TypeDish: convert struct to json err: %s", err)
		errors.HTTP.InternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Errorf("Handler: TypeDish: write response err: %s", err)
		return
	}
}
