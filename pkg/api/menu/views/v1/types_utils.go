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
