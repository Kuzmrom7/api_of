package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewAdress(obj *types.Adress) *Adress {
	return newAdress(obj)
}

func NewList(obj map[string]*types.Adress) *AdressList {
	if obj == nil {
		return nil
	}

	r := make(AdressList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}
