package place

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"strings"
)

type place struct {
	context context.Context
}

func New(c context.Context) *place {
	return &place{
		context: c,
	}
}

func (p *place) GetIDByName (name_typeplace string) (string, error){

	var (
		storage = ctx.Get().GetStorage()
	)

	typeplace_id, err := storage.Place().GetTypePlaceByName(p.context, strings.ToLower(name_typeplace))
	if err != nil {
		return "", err
	}
	if typeplace_id == "" {
		return "", nil
	}

	return typeplace_id, nil
}

func (p *place) Create(user, typeplace string, rq *request.RequestPlaceCreate) (*types.Place, error) {

	var (
		storage = ctx.Get().GetStorage()
		plc     = types.Place{}
	)

	plc.Meta.Adress = rq.Adress
	plc.Meta.City = rq.City
	plc.Meta.Name = rq.Name
	plc.Meta.Phone = rq.Phone
	plc.Meta.Url = rq.Url
	plc.Meta.TypePlaceID = typeplace
	plc.Meta.UserID = user

	if err := storage.Place().CreatePlace(p.context, &plc); err != nil {
		return nil, err
	}

	return &plc, nil
}

func (r *place) List() (map[string]*types.TypePlaces, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Place().List(r.context)
	if err != nil {
		return nil, err
	}
	return list, nil
}

