package pgsql

import (
	"fmt"
	"github.com/orderfood/api_of/pkg/storage"
	"encoding/json"
	"io"
	"log"
)
const (
	sqlGetUsers = `
		SELECT * FROM users
	`

	sqlCreqteUser = `
		INSERT INTO users (username, email) VALUES ($1, $2)
	`
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

type UserModel struct {

}

func GetUser() ([]byte, error) {
	log.Println("STORAGE--- GetUser()")
	rows, err := storage.DB.Query(sqlGetUsers)
	if err != nil {
		fmt.Println(err)
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

	productsJson, err := json.Marshal(users)

	return productsJson, nil
}

func CreateUser (productJson io.Reader) (int64,error) {
	log.Println("STORAGE--- CreateUser()")

	decoder := json.NewDecoder(productJson)
	user := User{}

	err := decoder.Decode(&user)
	if err != nil {
		log.Println(err)

	}

	result, err := storage.DB.Exec(sqlCreqteUser, user.Username, user.Email)

	if err != nil{
		log.Println(err)
	}

	lastInsertId,err := result.LastInsertId()

	return lastInsertId, nil
}