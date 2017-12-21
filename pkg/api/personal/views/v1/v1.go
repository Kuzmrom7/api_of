package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewPersonal(obj *types.Personal) *Personal{
	return newPersonal(obj)
}

func NewList(obj map[string]*types.TypePersonals) *TypePersonalList {
	if obj == nil {
		return nil
	}

	r := make(TypePersonalList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}