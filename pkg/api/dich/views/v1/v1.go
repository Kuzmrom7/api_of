package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewDich(obj *types.Dish) *Dich {
	return newDich(obj)
}

func NewList(obj map[string]*types.Dish) *DichList {
	if obj == nil {
		return nil
	}

	r := make(DichList, 0)
	for _, v := range obj {
		r = append(r, New(v))
	}
	return &r
}

func NewListType(obj map[string]*types.TypeDishes) *TypeDishList {
	if obj == nil {
		return nil
	}

	r := make(TypeDishList, 0)
	for _, v := range obj {
		r = append(r, Newt(v))
	}
	return &r
}
