package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewMenu(obj *types.Menu) *Menu {
	return newMenu(obj)
}

func NewList(obj map[string]*types.Menu) *MenuList {
	if obj == nil {
		return nil
	}

	r := make(MenuList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}
