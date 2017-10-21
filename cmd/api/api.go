package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"time"
)

var handler = http.HandlerFunc( func (w http.ResponseWriter, r *http.Request){

	w.Write([]byte("Hello, please change correct url!\n "))
})

func main(){
	r := mux.NewRouter()


	//ROUTES
	r.Handle("/", handler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}