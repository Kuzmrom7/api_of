package place

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"

)

type place struct {
	context context.Context
}

func New(c context.Context) *place{
	return &place{
		context: c,
	}
}

//func (u *place) GetByID (id string) (*types.User, error){
//	var (
//		storage = ctx.Get().GetStorage()
//	)
//
//	usr, err := storage.User().GetUserByID(u.context, id)
//	if err != nil {
//		return nil, err
//	}
//	if usr == nil {
//		return nil, nil
//	}
//
//	return usr, nil
//}

func (p *place) Create(user, typeplace string ,rq *request.RequestPlaceCreate) (*types.Place, error){

	var (
		storage = ctx.Get().GetStorage()
		plc = types.Place{}
	)

	plc.Meta.Adress = rq.Adress
	plc.Meta.City = rq.City
	plc.Meta.Name = rq.Name
	plc.Meta.Phone = rq.Phone
	plc.Meta.Url = rq.Url
	plc.Meta.TypePlaceID = typeplace
	plc.Meta.UserID = user

	//if err := p.CreatePlace(usr, *rq.Password); 	err != nil {
	//	return nil, err
	//}

	if err := storage.Place().CreatePlace(p.context, &plc); err != nil{
		return nil, err
	}

	return &plc, nil
}
