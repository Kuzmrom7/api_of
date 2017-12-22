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

func (p *dish) Create(rq *request.RequestDichCreate, typedishID string) (*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
		di      = types.Dish{}
	)

	di.Meta.Name = rq.Name
	di.Meta.Desc = rq.Desc
	di.Meta.Timemin = rq.Timemin
	di.Meta.TypeDishID = typedishID
	di.Meta.Url = rq.Url

	if err := storage.Dish().CreateDich(p.context, &di); err != nil {
		return nil, err
	}

	return &di, nil
}

func (p *dish) GetIDByName(name_dich string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	place_id, err := storage.Dish().GetIDdichByName(p.context, name_dich)
	if err != nil {
		return "", err
	}
	if place_id == "" {
		return "", nil
	}

	return place_id, nil
}

func (d *dish) Remove(id_dich string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	err := storage.Dish().Remove(d.context, id_dich)
	if err != nil {
		return err
	}
	return nil
}

func (r *dish) List() (map[string]*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Dish().List(r.context)
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