package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"log"
	"github.com/orderfood/api_of/pkg/storage/store"
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"errors"
)
const (
	sqlstrUserExistsByLogin = `

	`

	sqlGetUsers = `
		SELECT * FROM users
	`

	sqlCreqteUser = `
		INSERT INTO users (username, email, gravatar, password, salt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id;
	`
)

type UserStorage struct {
	storage.User
	client store.IDB
}

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

func (s *UserStorage) CheckExistsByLogin(ctx context.Context, login string) (bool, error) {
	result, err := s.client.Exec(sqlstrUserExistsByLogin, login)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return rows != 0, nil
}

//func GetUser() ([]byte, error) {
//	log.Println("STORAGE--- GetUser()")
//	rows, err := storage.DB.Query(sqlGetUsers)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	defer rows.Close()
//
//	users := make([]*Users, 0)
//
//	for rows.Next() {
//		us := new(Users)
//		err = rows.Scan(&us.User_id, &us.Username, &us.Email, &us.Created, &us.Updated )
//		if err != nil {
//			panic(err)
//		}
//		users = append(users, us)
//	}
//
//	productsJson, err := json.Marshal(users)
//
//	return productsJson, nil
//}

func (s *UserStorage) CreateUser (ctx context.Context, user *types.User) error {
	log.Println("STORAGE--- CreateUser()")

	if user == nil {
		err := errors.New("user can not be nil")
		return  err
	}

	var (
		err error
		id  store.NullString
	)

	tx, err := s.client.Begin()
	if err != nil {
		return err
	}

	tx.QueryRow(sqlCreqteUser, user.Meta.Username, user.Meta.Email, user.Meta.Gravatar,
		user.Security.Pass.Password, user.Security.Pass.Salt).Scan(&id)

	user.Meta.ID = id.String

	return err
}

func newUserStorage(client store.IDB) *UserStorage {
	s := new(UserStorage)
	s.client = client
	return s
}
