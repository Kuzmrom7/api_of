package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newPlace(obj *types.Place) *Place{
	u := new(Place)
	u.City = obj.Meta.City
	u.Name = obj.Meta.Name
	u.Url = obj.Meta.Url
	u.Phone = obj.Meta.Phone
	u.Adress = obj.Meta.Adress
	u.NameTypePlace = obj.Meta.TypePlaceID
	return u
}

func (obj *Place) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}