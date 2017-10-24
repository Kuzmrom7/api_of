package routes

import (
	"net/http"
	"encoding/json"

	"github.com/orderfood/api_of/pkg/storage"

)
type Users struct {
	Id        string `json:id`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func GetUser (w http.ResponseWriter, r *http.Request){

	rows, err := storage.DB.Query("SELECT * FROM users")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	defer rows.Close()

	users := make([]*Users, 0)

	for rows.Next() {
		us := new(Users)
		err = rows.Scan(&us.Id, &us.Name, &us.FirstName, &us.LastName)
		if err != nil {
			panic(err)
		}
		users = append(users, us)
	}

	if err != nil {
		panic(err)
	}

	productsJson, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(productsJson)
	w.Write([]byte("Hello"))
}
