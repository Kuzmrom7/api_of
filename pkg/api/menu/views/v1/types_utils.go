package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
)

func newMenu(obj *types.Menu) *Menu {
	m := new(Menu)

	m.Id = obj.Meta.ID
	m.Name = obj.Meta.Name
	m.Url = obj.Meta.Url

	return m
}

func (obj *Menu) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.Menu) *Menu {
	i := new(Menu)
	i.Id = obj.Meta.ID
	i.Name = obj.Meta.Name
	i.Url = obj.Meta.Url

	return i
}

func (obj *MenuList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &MenuList{}
	}
	return json.Marshal(obj)
}
