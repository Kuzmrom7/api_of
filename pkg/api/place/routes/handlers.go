package routes

import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/api/place/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/api/place"
)

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
	//fmt.Print(usrid1)
	//usrid := "42f39524-7cc3-4858-affb-a1c8822852bc"

	p := place.New(r.Context())

	typeplace_id, err := p.GetIDByName(rq.NameTypePlace)
	if err != nil {
		errors.HTTP.InternalServerError(w)
		return
	}
	if typeplace_id == "" {
		errors.New("type_place").NotFound().Http(w)
	}

	//TODO idTypePlace нужно запросом в базку получить по имени который в jsone приходит, те нужно сначала метод который получает id из базки
	//пока mock ресторана
	//var idTypePlace = "68c65b87-925b-4227-bada-c543b55048e2"

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

func PlaceList(w http.ResponseWriter, r *http.Request) {

	//if r.Context().Value("uid") == nil {
	//	errors.HTTP.Unauthorized(w)
	//	return
	//}

	items, err := place.New(r.Context()).List()
	if err != nil {
		log.Print("---------------")
		errors.HTTP.InternalServerError(w)
		return
	}

	response, err := v1.NewList(items).ToJson()
	if err != nil {
		log.Print("////////////////")
		errors.HTTP.InternalServerError(w)
	}

	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil {
		log.Println("Dich list response error")
		return
	}
}
