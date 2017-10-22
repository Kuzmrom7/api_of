package main

import (
	"github.com/gorilla/mux"
	"github.com/orderfood/api_of/pkg/util/http"


	users "github.com/orderfood/api_of/pkg/users/routes"
	"fmt"

	"log"
)

const (
	host = "localhost"
	port = 8080
)

var Routes = make([]http.Route, 0)

func AddRoutes(r ...[]http.Route){
	for i := range r{
		Routes = append(Routes, r[i]...)
		log.Print(Routes)
	}
}
func init(){
	//Здесь инициализация роутов
	//пример  AddRoutes(users.Routes)
	AddRoutes(users.Routes)
}

func main(){
	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(http.Headers)

	for _,route := range Routes {
		fmt.Print(r.Handle(route.Path, http.Handle(route.Handler, route.Middleware...)).Methods(route.Method))
	}

	http.Listen(host,port, r)
}

