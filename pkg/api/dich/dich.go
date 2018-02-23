package dich

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/common/types"
)

type dish struct {
	context context.Context
}

func New(c context.Context) *dish {
	return &dish{
		context: c,
	}
}

func (p *dish) Create(rq *request.DishCreate, userid string) (*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
		di      = types.Dish{}
	)

	log.Debugf("Dish: insert dish %#v", rq)

	di.Meta.Name = rq.Name
	di.Meta.Desc = rq.Desc
	di.Meta.Timemin = rq.Timemin
	di.Meta.TypeDishID = rq.IdTypeDish

	for _, u := range rq.Urls {
		di.Urls = append(di.Urls, types.UrlsOpt{
			Url: u.Url,
		})
	}

	di.Meta.UserID = userid

	if err := storage.Dish().CreateDish(p.context, &di); err != nil {
		log.Errorf("Dish: insert dish err: %s", err)
		return nil, err
	}

	return &di, nil
}

func (d *dish) Remove(id string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Dish: delete dish by id %s ", id)

	err := storage.Dish().RemoveDish(d.context, id)
	if err != nil {
		log.Errorf("Dish: delete dish by id err: %s", err)
		return err
	}
	return nil
}

func (r *dish) List(userid string) ([]*types.Dish, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Dish: list dishes")

	list, err := storage.Dish().List(r.context, userid)
	if err != nil {
		log.Errorf("Dish: list dishes err: %s", err)
		return nil, err
	}
	return list, nil
}

func (r *dish) TypeList() (map[string]*types.TypeDishes, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Dish: list type dish")

	list, err := storage.Dish().TypeList(r.context)
	if err != nil {
		log.Errorf("Dish: list type dish err: %s", err)
		return nil, err
	}
	return list, nil
}


func (u *dish) GetDishById(id string) (*types.Dish, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Dish: get dish by id %s", id)

	dish, err := storage.Dish().Fetch(u.context, id)
	if err != nil {
		log.Errorf("Dish: get dish by id `%s` err: %s", id, err)
		return nil, err
	}
	if dish == nil {
		log.Warnf("Dish: Dish by id `%s` not found", id)
		return nil, nil
	}

	return dish, nil
}

func (p *dish) Update(usrid string, rq *request.DishUpdate, dish *types.Dish) error {
	var (
		err     error
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Dish: update dish %#v", rq)

	dish.Meta.ID = rq.Id

	if rq.Timemin != nil {
		dish.Meta.Timemin = *rq.Timemin
	}
	if rq.Desc != nil {
		dish.Meta.Desc = *rq.Desc
	}

	if err = storage.Dish().Update(p.context, dish); err != nil {
		log.Errorf("Dish: update dish err: %s", err)
		return err
	}
	return nil
}
