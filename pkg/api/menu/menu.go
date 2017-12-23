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

func (p *menu) GetIDplaceByName(name_place string) (string, error) {

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

func (p *menu) GetIDMenuByName(name_menu string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	menu_id, err := storage.Menu().GetIDmenuByName(p.context, name_menu)
	if err != nil {
		return "", err
	}
	if menu_id == "" {
		return "", nil
	}

	return menu_id, nil
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

func (r *menu) List(placeid string) (map[string]*types.Menu, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Menu().List(r.context, placeid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (p *menu) CreateMenuDish(menuid, dishid string)  error {

	var (
		storage = ctx.Get().GetStorage()

	)

	if err := storage.Menu().CreateMenuDish(p.context, menuid, dishid); err != nil {
		return  err
	}

	return  nil
}

func (u *menu) GetMenuByIDPlaceAndNameMenu(idplace, namem string) (*types.Menu, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	plc, err := storage.Menu().Fetch(u.context, idplace, namem)
	if err != nil {
		return nil, err
	}
	if plc == nil {
		return nil, nil
	}

	return plc, nil
}