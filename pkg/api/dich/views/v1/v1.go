package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewDich(obj *types.Dich) *Dich {
	return newDich(obj)
}

func NewList(obj map[string]*types.Dich) *DichList {
	if obj == nil {
		return nil
	}

	r := make(DichList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}
