package menu

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/menu/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
)

type menu struct {
	context context.Context
}

func New(c context.Context) *menu {
	return &menu{
		context: c,
	}
}

func (p *menu) Create(rq *request.MenuCreate) (*types.Menu, error) {

	var (
		storage = ctx.Get().GetStorage()
		men     = types.Menu{}
	)

	log.Debug("Menu: insert menu ")

	men.Meta.Name = rq.Name
	men.Meta.PlaceID = rq.Id_place
	men.Meta.Url = rq.Url

	if err := storage.Menu().CreateMenu(p.context, &men); err != nil {
		log.Errorf("Menu: insert menu err: %s", err)
		return nil, err
	}

	return &men, nil
}

func (r *menu) List(placeid string) (map[string]*types.Menu, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Menu: List: get list menu by place id %s", placeid)

	list, err := storage.Menu().List(r.context, placeid)
	if err != nil {
		log.Errorf("Menu: List: get list menu by place id err: %s", err)
		return nil, err
	}
	return list, nil
}

func (p *menu) CreateMenuDish(menuid, dishid string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Menu: Dish: add dish by id %s in menu by id %s", menuid, dishid)

	if err := storage.Menu().InsertDishInMenu(p.context, menuid, dishid); err != nil {
		log.Errorf("Menu: Dish: add dish in menu err: %s", err)
		return err
	}

	return nil
}

func (p *menu) CheckUniqueDishInMenu(menuid, dishid string) (bool, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Menu: Dish: check unique add dish by id %s in menu by id %s", menuid, dishid)

	exists, err := storage.Menu().CheckUnique(p.context, menuid, dishid)
	if err != nil {
		log.Errorf("Menu: Dish: check unique add dish in menu err: %s", err)
		return false, err
	}

	return exists, nil
}

func (p *menu) RemoveMenuDish(menuid, dishid string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Menu: Dish: delete dish by id %s from menu by id %s", menuid, dishid)

	if err := storage.Menu().DeleteDishInMenu(p.context, menuid, dishid); err != nil {
		log.Errorf("Menu: Dish: delete dish from menu err: %s", err)
		return err
	}

	return nil
}

func (u *menu) GetMenuByID(id string) (*types.Menu, error) {
	var (
		storage = ctx.Get().GetStorage()
	)
	log.Debugf("Menu: Info: get menu by id %s", id)

	men, err := storage.Menu().Fetch(u.context, id)
	if err != nil {
		log.Errorf("Menu: get menu by id `%s` err: %s", id, err)
		return nil, err
	}
	if men == nil {
		log.Warnf("Menu: menu by id `%s` not found", id)
		return nil, nil
	}

	return men, nil
}

func (r *menu) ListMenuDish(menuid, placeid string) ([]*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Menu: Dish: List: list dish in menu")

	list, err := storage.Menu().ListDishesInMenu(r.context, menuid, placeid)
	if err != nil {
		log.Errorf("Menu: Dish: List: list dish in menu err: %s", err)
		return nil, err
	}
	return list, nil
}

func (r *menu) ListDishNotMenu(menuid, placeid string) ([]*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Menu: Dish: List: list dish not in menu")

	list, err := storage.Menu().ListDishesNotMenu(r.context, menuid, placeid)
	if err != nil {
		log.Errorf("Menu: Dish: List: list dish not in menu err: %s", err)
		return nil, err
	}
	return list, nil
}
