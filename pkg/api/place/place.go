package place

import (
	"context"

	//ctx "github.com/orderfood/api_of/pkg/api/context"
)

type place struct {
	context context.Context
}

func New(c context.Context) *place {
	return &place{
		context: c,
	}
}

