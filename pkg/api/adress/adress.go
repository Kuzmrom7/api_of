package adress

import (
	"github.com/orderfood/api_of/pkg/log"
	"github.com/orderfood/api_of/pkg/api/adress/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
	ctx "github.com/orderfood/api_of/pkg/api/context"
	"context"
)

type adress struct {
	context context.Context
}

func New(c context.Context) *adress {
	return &adress{
		context: c,
	}
}

func (p *adress) Create(rq *request.AdressCreate) (*types.Adress, error) {

	var (
		storage = ctx.Get().GetStorage()
		adr     = types.Adress{}
	)

	log.Debug("Adress: insert adress ")

	adr.Meta.Name = rq.Name
	adr.Meta.PlaceID = rq.Id_place

	if err := storage.Adress().CreateAdress(p.context, &adr); err != nil {
		log.Errorf("Adress: insert adress err: %s", err)
		return nil, err
	}

	return &adr, nil
}

func (r *adress) List(placeid string) (map[string]*types.Adress, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	log.Debugf("Adress: List: get list adress by place id %s", placeid)

	list, err := storage.Adress().List(r.context, placeid)
	if err != nil {
		log.Errorf("Adress: List: get list adress by place id err: %s", err)
		return nil, err
	}
	return list, nil
}
