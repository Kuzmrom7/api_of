package routes

import (
	"net/http"
	"log"

	"github.com/orderfood/api_of/pkg/storage/pgsql"
)



func GetUser (w http.ResponseWriter, r *http.Request){

	productsJson, err := pgsql.GetUser()

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)
}

func UserCreate (w http.ResponseWriter, r *http.Request) {

	res, err := pgsql.CreateUser(r.Body)
	if err != nil {
		log.Println(err)
	}

	log.Println("Create user id: " , res)
	w.WriteHeader(http.StatusOK)

}