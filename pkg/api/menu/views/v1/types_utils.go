package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newMenu(obj *types.Menu) *Menu {
	m := new(Menu)
	m.Name = obj.Meta.Name
	m.Url = obj.Meta.Url
	m.Created = obj.Meta.Created
	m.Updated = obj.Meta.Updated
	return m
}

func (obj *Menu) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}


func New(obj *types.Menu) *Menu {
	i := new(Menu)
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