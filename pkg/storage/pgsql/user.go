package pgsql

import (
	"github.com/orderfood/api_of/pkg/storage/storage"
	"log"
	"github.com/orderfood/api_of/pkg/storage/store"
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"errors"
	"database/sql"

)
const (

	sqlstrUserGetByLogin = ``

	sqlstrUserExistsByLogin = `
		SELECT TRUE
		FROM users
		WHERE users.username = $1 OR users.email = $1
	`

	sqlstrUserGetById = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.user_id = $1;
	`

	sqlCreateUser = `
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
	Username      string `json:"username"`
	Email 		  string `json:"email"`
}

type userModel struct {
	id 				store.NullString
	username 	store.NullString
	email 		store.NullString
	gravatar 	store.NullString
	password 	store.NullString
	salt 			store.NullString
}

func (um *userModel) convert() *types.User{
	var u = new(types.User)
	u.Meta.Username = um.username.String
	u.Meta.ID = um.id.String
	u.Meta.Email = um.email.String
	u.Meta.Gravatar = um.gravatar.String
	u.Security.Pass.Password = um.password.String
	u.Security.Pass.Salt = um.salt.String

	return u
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

func (s *UserStorage) GetByLogin (ctx context.Context, login string) (*types.User, error) {
	var (
		err error
		um = new(userModel)
	)

	err = s.client.QueryRow(sqlstrUserGetByLogin, login).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)

	switch err{
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	usr := um.convert()

	return usr, nil
}

func (s *UserStorage)GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var (
		err error
		um = new(userModel)
	)

	err = s.client.QueryRow(sqlstrUserGetById, id).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)
	switch err{
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	usr := um.convert()

	return usr, nil
}

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

	err = s.client.QueryRow( sqlCreateUser, user.Meta.Username, user.Meta.Email, user.Meta.Gravatar,
		user.Security.Pass.Password, user.Security.Pass.Salt).Scan(&id)

	user.Meta.ID = id.String

	return err
}

func newUserStorage(client store.IDB) *UserStorage {
	s := new(UserStorage)
	s.client = client
	return s
}
