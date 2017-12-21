package pgsql

import (
	"log"
	"github.com/orderfood/api_of/pkg/storage/store"
	"context"
	"github.com/orderfood/api_of/pkg/common/types"
	"errors"
	"database/sql"
)

func (um *userModel) convert() *types.User {
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

func (s *UserStorage) GetByLogin(ctx context.Context, login string) (*types.User, error) {
	var (
		err error
		um  = new(userModel)
	)

	err = s.client.QueryRow(sqlstrUserGetByLogin, login).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)

	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	usr := um.convert()

	return usr, nil
}

func (s *UserStorage) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var (
		err error
		um  = new(userModel)
	)

	err = s.client.QueryRow(sqlstrUserGetById, id).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)
	switch err {
	case nil:
	case sql.ErrNoRows:
		return nil, nil
	default:
		return nil, err
	}

	usr := um.convert()

	return usr, nil
}

func (s *UserStorage) CreateUser(ctx context.Context, user *types.User) error {
	log.Println("STORAGE--- CreateUser()")

	if user == nil {
		err := errors.New("user can not be nil")
		return err
	}

	var (
		err error
		id  store.NullString
	)

	err = s.client.QueryRow(sqlCreateUser, user.Meta.Username, user.Meta.Email, user.Meta.Gravatar,
		user.Security.Pass.Password, user.Security.Pass.Salt).Scan(&id)

	user.Meta.ID = id.String

	return err
}

