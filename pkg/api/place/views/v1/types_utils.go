package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newPlace(obj *types.Place) *Place {
	p := new(Place)
	p.City = obj.Meta.City
	p.Name = obj.Meta.Name
	p.Url = obj.Meta.Url
	p.Phone = obj.Meta.Phone
	p.Id = obj.Meta.ID

	if obj.Adresses != nil {
		for i, adr := range obj.Adresses {
			p.Adresses[i].Adress = adr.Adress
		}
	}

	if obj.TypesPlace != nil {
		for i, typepl := range obj.TypesPlace {
			p.TypesPlace[i].ID = typepl.ID
			p.TypesPlace[i].NameType = typepl.NameType
		}
	}

	return p
}

func (obj *Place) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.TypePlaces) *TypePlace {
	i := new(TypePlace)
	i.Meta.Name = obj.NameType
	i.Meta.ID = obj.ID
	return i
}

func (obj *TypePlaceList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &TypePlaceList{}
	}
	return json.Marshal(obj)
}
