package place

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
)

type place struct {
	context context.Context
}

func New(c context.Context) *place {
	return &place{
		context: c,
	}
}

func (p *place) Create(user string, rq *request.PlaceCreate) (*types.Place, error) {

	var (
		storage = ctx.Get().GetStorage()
		plc     = types.Place{}
	)

	log.Debugf("Place: create place %#v", rq)

	plc.Meta.Name = rq.Name

	for _, typepl := range rq.TypesPlace {
		plc.TypesPlace = append(plc.TypesPlace, types.TypePlaces{
			ID:       typepl.IdTypePlace,
			NameType: typepl.NameTypePlace,
		})
	}

	plc.Meta.UserID = user

	if err := storage.Place().CreatePlace(p.context, &plc); err != nil {
		log.Errorf("Place: insert place err: %s", err)
		return nil, err
	}

	return &plc, nil
}

func (r *place) ListType() (map[string]*types.TypePlaces, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Place: list type place")

	list, err := storage.Place().ListType(r.context)
	if err != nil {
		log.Errorf("Place: list type place err: %s", err)
		return nil, err
	}
	return list, nil
}

func (r *place) List() ([]*types.Place, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debug("Place: list places")

	list, err := storage.Place().List(r.context)
	if err != nil {
		log.Errorf("Place: list places err: %s", err)
		return nil, err
	}
	return list, nil
}

func (u *place) GetPlaceByID(id string) (*types.Place, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Place: get place by id %s", id)

	plc, err := storage.Place().GetPlaceByID(u.context, id)
	if err != nil {
		log.Errorf("Place: get place by id `%s` err: %s", id, err)
		return nil, err
	}
	if plc == nil {
		log.Warnf("Place: Place by id `%s` not found", id)
		return nil, nil
	}

	return plc, nil
}

func (u *place) GetPlaceByIDUsr(id string) (*types.Place, error) {
	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Place: get place by user id %s", id)

	plc, err := storage.Place().GetPlaceByIDUser(u.context, id)
	if err != nil {
		log.Errorf("Place: get place by user id `%s` err: %s", id, err)
		return nil, err
	}
	if plc == nil {
		log.Warnf("Place: Place by user id `%s` not found", id)
		return nil, nil
	}

	return plc, nil
}

func (p *place) Update(place *types.Place, rq *request.PlaceUpdate) error {
	var (
		err     error
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Place: update place %#v", rq)

	if rq.Url != nil {
		place.Meta.Url = *rq.Url
	}
	if rq.City != nil {
		place.Meta.City = *rq.City
	}

	if rq.Phone != nil {
		place.Meta.Phone = *rq.Phone
	}

	if rq.Logo != nil {
		place.Meta.Logo = *rq.Logo
	}

	place.Meta.ID = rq.Id

	if err = storage.Place().Update(p.context, place); err != nil {
		log.Errorf("Place: update place err: %s", err)
		return err
	}
	return nil
}
