package dich

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"strings"
)

type dish struct {
	context context.Context
}

func New(c context.Context) *dish {
	return &dish{
		context: c,
	}
}

func (p *dish) Create(rq *request.RequestDichCreate, typedishID, userid string) (*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
		di      = types.Dish{}
	)

	di.Meta.Name = rq.Name
	di.Meta.Desc = rq.Desc
	di.Meta.Timemin = rq.Timemin
	di.Meta.TypeDishID = typedishID
	di.Meta.Url = rq.Url
	di.Meta.UserID = userid

	if err := storage.Dish().CreateDish(p.context, &di); err != nil {
		return nil, err
	}

	return &di, nil
}

func (p *dish) GetIDdishByName(name_dich, usrid string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	dish_id, err := storage.Dish().GetIdDishByName(p.context, name_dich, usrid)
	if err != nil {
		return "", err
	}
	if dish_id == "" {
		return "", nil
	}

	return dish_id, nil
}

func (d *dish) Remove(id_dich string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	err := storage.Dish().RemoveDish(d.context, id_dich)
	if err != nil {
		return err
	}
	return nil
}

func (r *dish) List(userid string) (map[string]*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Dish().List(r.context, userid)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *dish) TypeList() (map[string]*types.TypeDishes, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Dish().TypeList(r.context)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (p *dish) GetIDTypeDishByName(type_name string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	type_name = strings.ToLower(type_name)

	typedish_id, err := storage.Dish().GetTypeDishIDByName(p.context, type_name)
	if err != nil {
		return "", err
	}
	if typedish_id == "" {
		return "", nil
	}

	return typedish_id, nil
}

func (u *dish) GetDishByUsrIdAndDishName(usrid, name string) (*types.Dish, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	dish, err := storage.Dish().Fetch(u.context, usrid, name)
	if err != nil {
		return nil, err
	}
	if dish == nil {
		return nil, nil
	}

	return dish, nil
}


func (p *dish) Update(usrid string, rq *request.RequestDichUpdate, dish *types.Dish) error{
	var (
		err     error
		storage = ctx.Get().GetStorage()
	)

	if rq.Timemin != nil {
		dish.Meta.Timemin = *rq.Timemin
	}
	if rq.Desc != nil {
		dish.Meta.Desc = *rq.Desc
	}

	if err = storage.Dish().Update(p.context, usrid, dish); err != nil {
		return err
	}
	return nil
}
