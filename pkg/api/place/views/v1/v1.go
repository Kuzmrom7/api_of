package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewPlace(obj *types.Place) *Place {
	return newPlace(obj)
}

func NewListType(obj map[string]*types.TypePlaces) *TypePlaceList {
	if obj == nil {
		return nil
	}

	r := make(TypePlaceList, 0)
	for _, v := range obj {
		r = append(r, NewT(v))
	}
	return &r
}

func NewList(obj []*types.Place) *PlaceList {
	if obj == nil {
		return nil
	}

	r := make(PlaceList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}
