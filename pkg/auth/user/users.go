package user

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/auth/user/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/util/generator"
	"github.com/orderfood/api_of/pkg/log"
)

type user struct {
	context context.Context
}

func New(c context.Context) *user {
	return &user{
		context: c,
	}
}

//-----------------------------------------------------------------------------------------------

func (u *user) CheckExists(login string) (bool, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("User: check exists user %s", login)

	exists, err := storage.User().CheckExistsByLogin(u.context, login)
	if err != nil {
		log.Errorf("User: check exists `%s` err: %s", login, err)
		return false, err
	}

	return exists, nil
}

func (u *user) GetByID(id string) (*types.User, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("User: get user by id %s", id)

	usr, err := storage.User().GetUserByID(u.context, id)
	if err != nil {
		log.Errorf("User: get user by id err %s", err)
		return nil, err
	}
	if usr == nil {
		return nil, nil
	}

	return usr, nil
}

func (u *user) Create(rq *request.RequestUserCreate) (*types.User, error) {

	var (
		storage = ctx.Get().GetStorage()
		usr     = new(types.User)
	)

	log.Debugf("User: create user email: %s, username: %s", *rq.Email, *rq.Username)

	usr.Meta.Gravatar = generator.GenerateGravatar(*rq.Email)
	usr.Meta.Username = *rq.Username
	usr.Meta.Email = *rq.Email

	salt, err := generator.GenerateSalt(*rq.Password)
	if err != nil {
		log.Errorf("User: generate salt err: %s", err)
		return nil, err
	}

	pass, err := generator.Generatepassword(*rq.Password, salt)
	if err != nil {
		log.Errorf("User: generate password err: %s", err)
		return nil, err
	}

	usr.Security.Pass.Password = pass
	usr.Security.Pass.Salt = salt

	if err := storage.User().CreateUser(u.context, usr); err != nil {
		log.Errorf("User: insert account err: %s", err)
		return nil, err
	}

	return usr, nil
}
