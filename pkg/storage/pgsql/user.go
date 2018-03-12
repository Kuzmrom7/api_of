package pgsql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/storage/store"
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
	log.Debugf("Storage: User: Exists: check account exists by: %s", login)

	const sqlstrUserExistsByLogin = `
		SELECT TRUE
		FROM users
		WHERE users.username = $1 OR users.email = $1
	`

	result, err := s.client.Exec(sqlstrUserExistsByLogin, login)
	if err != nil {
		log.Errorf("Storage: User: Exists: find user query err: %s", err)
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		log.Errorf("Storage: User: Exists: check query affected err: %s", err)
		return false, err
	}

	return rows != 0, nil
}

func (s *UserStorage) GetByLogin(ctx context.Context, login string) (*types.User, error) {

	log.Debugf("Storage: User: get user by login: %s", login)

	var (
		err error
		um  = new(userModel)
	)

	const sqlstrUserGetByLogin = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.username = $1;`

	err = s.client.QueryRow(sqlstrUserGetByLogin, login).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)
	if err != nil {
		log.Errorf("Storage: User: get user by login err: %s", err)
		return nil, err
	}

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

func (s *UserStorage) GetById(ctx context.Context, id string) (*types.User, error) {

	log.Debugf("Storage: User: get user by id: %s", id)

	var (
		err error
		um  = new(userModel)
	)

	const sqlstrUserGetId = `
		SELECT users.user_id, users.username, users.email, users.gravatar, users.password, users.salt
		FROM users
		WHERE users.user_id = $1;`

	err = s.client.QueryRow(sqlstrUserGetId, id).Scan(&um.id, &um.username, &um.email,
		&um.gravatar, &um.password, &um.salt)
	if err != nil {
		log.Errorf("Storage: User: get user by id err: %s", err)
		return nil, err
	}

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

	var (
		err error
		id  store.NullString
	)

	log.Debugf("Storage: User: insert user: %#v", user)

	if user == nil {
		err := errors.New("user can not be nil")
		log.Errorf("Storage: User: insert user err: %s", err)
		return err
	}

	const sqlCreateUser = `
		INSERT INTO users (username, email, gravatar, password, salt)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING user_id;
	`

	err = s.client.QueryRow(sqlCreateUser, user.Meta.Username, user.Meta.Email, user.Meta.Gravatar,
		user.Security.Pass.Password, user.Security.Pass.Salt).Scan(&id)
	if err != nil {
		log.Errorf("Storage: User: Insert: user insert err: %s", err)
		return err
	}

	user.Meta.ID = id.String

	return err
}
