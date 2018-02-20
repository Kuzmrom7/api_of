package place

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/place/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"github.com/orderfood/api_of/pkg/log"
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

func (p *place) GetIDTypePlaceByName(name_typeplace string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Place: get id type place by name %s", name_typeplace)

	name_typeplace = strings.ToLower(name_typeplace)

	typeplace_id, err := storage.Place().GetTypePlaceByName(p.context, name_typeplace)
	if err != nil {
		log.Errorf("Place: get id type place by name `%s` err: %s", name_typeplace, err)
		return "", err
	}
	if typeplace_id == "" {
		log.Warnf("Place: id type place by name `%s` not found", name_typeplace)
		return "", nil
	}

	return typeplace_id, nil
}

func (p *place) Create(user, typeplace string, rq *request.RequestPlaceCreate) (*types.Place, error) {

	var (
		storage = ctx.Get().GetStorage()
		plc     = types.Place{}
	)

	log.Debugf("Place: create place %#v", rq)

	plc.Meta.Adress = rq.Adress
	plc.Meta.City = rq.City
	plc.Meta.Name = rq.Name
	plc.Meta.Phone = rq.Phone
	plc.Meta.Url = rq.Url
	plc.Meta.TypePlaceID = typeplace
	plc.Meta.UserID = user

	if err := storage.Place().CreatePlace(p.context, &plc); err != nil {
		log.Errorf("Place: insert place err: %s", err)
		return nil, err
	}

	return &plc, nil
}

func (r *place) List() (map[string]*types.TypePlaces, error) {

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

func (p *place) Update(place *types.Place, rq *request.RequestPlaceUpdate) error {
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

	if rq.Adress != nil {
		place.Meta.Adress = *rq.Adress
	}

	if rq.Phone != nil {
		place.Meta.Phone = *rq.Phone
	}

	if err = storage.Place().Update(p.context, place); err != nil {
		log.Errorf("Place: update place err: %s", err)
		return err
	}
	return nil
}
