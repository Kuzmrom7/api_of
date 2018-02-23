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

	p.Adresses = make([]*AdressOpt, 0)
	if obj.Adresses != nil {
		for _, adr := range obj.Adresses {
			a := new(AdressOpt)

			a.Adress = adr.Adress

			p.Adresses = append(p.Adresses, a)
		}
	}
	p.TypesPlace = make([]*TypePlaces, 0)
	if obj.TypesPlace != nil {
		for _, typepl := range obj.TypesPlace {
			t := new(TypePlaces)

			t.ID = typepl.ID
			t.NameType = typepl.NameType

			p.TypesPlace = append(p.TypesPlace, t)
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
