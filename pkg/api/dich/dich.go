package dich

import (
	"context"

	ctx "github.com/orderfood/api_of/pkg/api/context"
	"github.com/orderfood/api_of/pkg/api/dich/routes/request"
	"github.com/orderfood/api_of/pkg/common/types"
)

type dich struct {
	context context.Context
}

func New(c context.Context) *dich {
	return &dich{
		context: c,
	}
}

func (p *dich) Create(rq *request.RequestDichCreate) (*types.Dich, error) {

	var (
		storage = ctx.Get().GetStorage()
		di      = types.Dich{}
	)

	di.Meta.Name = rq.Name
	di.Meta.Desc = rq.Desc
	di.Meta.Timemin = rq.Timemin

	if err := storage.Dich().CreateDich(p.context, &di); err != nil {
		return nil, err
	}

	return &di, nil
}

func (p *dich) GetIDByName(name_dich string) (string, error) {

	var (
		storage = ctx.Get().GetStorage()
	)

	place_id, err := storage.Dich().GetIDdichByName(p.context, name_dich)
	if err != nil {
		return "", err
	}
	if place_id == "" {
		return "", nil
	}

	return place_id, nil
}

func (d *dich) Remove(id_dich string) error {

	var (
		storage = ctx.Get().GetStorage()
	)

	err := storage.Dich().Remove(d.context, id_dich)
	if err != nil {
		return err
	}
	return nil
}
