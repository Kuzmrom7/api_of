package routes

import (
	"net/http"
	"log"

	//"github.com/orderfood/api_of/pkg/storage/pgsql"
	"github.com/orderfood/api_of/pkg/auth/user/routes/request"
	"github.com/orderfood/api_of/pkg/auth/user"
	//"github.com/orderfood/api_of/pkg/common/errors"
)



func GetUser (w http.ResponseWriter, r *http.Request){
	//
	//productsJson, err := pgsql.GetUser()
	//
	//if err != nil {
	//	log.Println(err)
	//	w.WriteHeader(http.StatusBadRequest)
	//}
	//
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(http.StatusOK)
	//w.Write(productsJson)
}

func UserCreate (w http.ResponseWriter, r *http.Request) {

	var (
		err error
	)

	rq := new(request.RequestUserCreate)
	if err := rq.DecodeAndValidate(r.Body); err != nil {
		err.Http(w)
		return
	}

	u := user.New(r.Context())
	//TODO on feature
	//exists, err := u.CheckExists(*rq.Username)
	//if err != nil && exists {
	//	errors.New("user").NotUnique("username").Http(w)
	//	return
	//}
	//if err != nil{
	//	errors.HTTP.InternalServerError(w)
	//}
	//log.Print("I here3")
	//exists, err = u.CheckExists(*rq.Email)
	//if err != nil && exists {
	//	errors.New("user").NotUnique("email").Http(w)
	//	return
	//}
	//if err != nil{
	//	errors.HTTP.InternalServerError(w)
	//}

	usr, err := u.Create(rq)

	//TODO Session
	log.Println("Create user id: " , usr.Meta.ID)
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(usr.Meta.ID)); err != nil{
		log.Println("User write response error")
		return
	}


}