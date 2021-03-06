package personal

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/personal/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	"strings"
)

type personal struct {
	context context.Context
}

func New(c context.Context) *personal {
	return &personal{
		context: c,
	}
}

func (p *personal) GetIDTypePersonByName(name_typeperson string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	name_typeperson = strings.ToLower(name_typeperson)

	typeplace_id, err := storage.Personal().GetTypePersonIDByName(p.context, name_typeperson)
	if err != nil {
		return "", err
	}
	if typeplace_id == "" {
		return "", nil
	}

	return typeplace_id, nil
}

//
func (p *personal) Create(typeperson, place string, rq *request.RequestPersonCreate) (*types.Personal, error) {

	var (
		storage = ctx.Get().GetStorage()
		pers    = types.Personal{}
	)

	pers.Meta.Fio = rq.Fio
	pers.Meta.Phone = rq.Phone
	pers.Meta.PlaceID = place
	pers.Meta.TypePersonalID = typeperson

	if err := storage.Personal().CreatePerson(p.context, &pers); err != nil {
		return nil, err
	}

	return &pers, nil
}

func (r *personal) ListType() (map[string]*types.TypePersonals, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Personal().ListType(r.context)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (r *personal) List(placeid string) (map[string]*types.Personal, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	list, err := storage.Personal().List(r.context, placeid)
	if err != nil {
		return nil, err
	}
	return list, nil
}
