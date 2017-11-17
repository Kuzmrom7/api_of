package user

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/auth/user/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"

	"github.com/orderfood/api_of/pkg/util/generator"
)

type user struct {
	context context.Context
}

func New(c context.Context) *user{
	return &user{
		context: c,
	}
}

//-----------------------------------------------------------------------------------------------

func (u *user) CheckExists(login string) (bool, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	exists, err := storage.User().CheckExistsByLogin(u.context, login)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (u *user) Create(rq *request.RequestUserCreate) (*types.User, error){

	var (
		storage = ctx.Get().GetStorage()
		usr = new(types.User)
	)

	usr.Meta.Gravatar = generator.GenerateGravatar(*rq.Email)
	usr.Meta.Username = *rq.Username
	usr.Meta.Email = *rq.Email

	if err := u.CreatePassword(usr, *rq.Password); 	err != nil {
		return nil, err
	}

	if err := storage.User().CreateUser(u.context, usr); err != nil{
		return nil, err
	}

//TODO Add registry with generate token jwt

	return usr, nil
}

func (u *user) CreatePassword(user *types.User, text string) error{

	salt, err := generator.GenerateSalt(text)
	if err != nil {
		return  err
	}

	pass, err := generator.Generatepassword(text, salt)
	if err != nil {
		return err
	}

	user.Security.Pass.Password = pass
	user.Security.Pass.Salt = salt

	return  nil
}