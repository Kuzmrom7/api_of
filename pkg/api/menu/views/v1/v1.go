package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewMenu(obj *types.Menu) *Menu {
	return newMenu(obj)
}
