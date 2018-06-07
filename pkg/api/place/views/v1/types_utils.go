package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
)

func newPlace(obj *types.Place) *Place {
	p := new(Place)
	p.City = obj.Meta.City
	p.Name = obj.Meta.Name
	p.Url = obj.Meta.Url
	p.Phone = obj.Meta.Phone
	p.Id = obj.Meta.ID
	p.Logo = obj.Meta.Logo

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

func NewT(obj *types.TypePlaces) *TypePlace {
	i := new(TypePlace)
	i.Meta.Name = obj.NameType
	i.Meta.ID = obj.ID
	return i
}

func New(obj *types.Place) *Place {
	i := new(Place)
	i.Id = obj.Meta.ID
	i.Url = obj.Meta.Url
	i.Name = obj.Meta.Name
	i.City = obj.Meta.City
	i.Phone = obj.Meta.Phone
	i.Logo = obj.Meta.Logo

	i.TypesPlace = make([]*TypePlaces, 0)
	if obj.TypesPlace != nil {
		for _, typepl := range obj.TypesPlace {
			t := new(TypePlaces)

			t.ID = typepl.ID
			t.NameType = typepl.NameType

			i.TypesPlace = append(i.TypesPlace, t)
		}
	}

	return i
}

func (obj *TypePlaceList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &TypePlaceList{}
	}
	return json.Marshal(obj)
}

func (obj *PlaceList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &PlaceList{}
	}
	return json.Marshal(obj)
}
