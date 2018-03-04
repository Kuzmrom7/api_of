package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newAdress(obj *types.Adress) *Adress {
	m := new(Adress)

	m.Id = obj.Meta.ID
	m.Name = obj.Meta.Name
	m.PlaceID = obj.Meta.PlaceID

	return m
}

func (obj *Adress) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.Adress) *Adress {
	i := new(Adress)

	i.Id = obj.Meta.ID
	i.Name = obj.Meta.Name
	i.PlaceID = obj.Meta.PlaceID

	return i
}

func (obj *AdressList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &AdressList{}
	}
	return json.Marshal(obj)
}
