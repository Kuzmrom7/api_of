package routes

import (
	"net/http"
	"encoding/json"

	"github.com/orderfood/api_of/pkg/storage"

	"fmt"
	"log"
)

type Users struct {
	User_id   	  string `json:"user_id"`
	Username      string `json:"username"`
	Email 		  string `json:"email"`
	Created       string `json:"created"`
	Updated		  string `json:"updated"`
}

type User struct {
	/*User_id   	  string `json:"user_id"`*/
	Username      string `json:"username"`
	Email 		  string `json:"email"`
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
		err = rows.Scan(&us.User_id, &us.Username, &us.Email, &us.Created, &us.Updated )
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

func UserCreate (w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	user := User{}

	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := storage.DB.Exec("INSERT INTO users (username, email) VALUES ($1, $2)", user.Username, user.Email)
	if err != nil{
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	lastInsertId,err := result.LastInsertId()
	fmt.Println(lastInsertId)
	w.WriteHeader(http.StatusOK)

}