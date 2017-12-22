package menu

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/menu/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
)

type menu struct {
	context context.Context
}

func New(c context.Context) *menu {
	return &menu{
		context: c,
	}
}

func (p *menu) GetIDByName(name_place string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	place_id, err := storage.Menu().GetPlaceByName(p.context, name_place)
	if err != nil {
		return "", err
	}
	if place_id == "" {
		return "", nil
	}

	return place_id, nil
}

//
func (p *menu) Create(place string, rq *request.RequestMenuCreate) (*types.Menu, error) {

	var (
		storage = ctx.Get().GetStorage()
		men     = types.Menu{}
	)

	men.Meta.Name = rq.Name
	men.Meta.PlaceID = place
	men.Meta.Url = rq.Url

	if err := storage.Menu().CreateMenu(p.context, &men); err != nil {
		return nil, err
	}

	return &men, nil
}
