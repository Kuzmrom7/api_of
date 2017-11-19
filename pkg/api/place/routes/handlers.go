package routes

import (
	"net/http"


	"github.com/orderfood/api_of/pkg/common/errors"
)

func GetPlace(w http.ResponseWriter, r *http.Request){
	if r.Context().Value("uid") == nil {
		errors.HTTP.Unauthorized(w)
		return
	}

	var (
	/*	err =     error
		id = r.Context().Value("uid").(string)*/
	)
}

func CreatePlace(){

}

func UpdatePlace(){

}