package routes

import (
	"net/http"
	"log"

	//sv1 "github.com/orderfood/api_of/pkg/api/place/views/v1"
	//"github.com/orderfood/api_of/pkg/auth/user/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/views/v1"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/common/errors"

	"github.com/orderfood/api_of/pkg/api/place"
)

func PlaceCreate (w http.ResponseWriter, r *http.Request) {

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

	usr := r.Context().Value("uid").(*types.User)

	p := place.New(r.Context())

	//TODO idTypePlace нужно запросом в базку получить по имени который в jsone приходит, те нужно сначала метод который получает id из базки
	//пока mock ресторана
	var idTypePlace = "68c65b87-925b-4227-bada-c543b55048e2"

	plc, err := p.Create(usr.Meta.ID, idTypePlace ,rq)

	response, err := v1.NewPlace(plc).ToJson()
	if err != nil {
		errors.HTTP.InternalServerError(w)
	}

	///log.Println("Create user id: " , usr.Meta.ID, " username: " , usr.Meta.Username)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(response); err != nil{
		log.Println("Place write response error")
		return
	}
}
