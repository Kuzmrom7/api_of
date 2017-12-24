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

func NewListD(dish map[string]*types.Dish, typedish map[string]*types.TypeDishes) *TypeDishListinMenu {
	if dish == nil {
		return nil
	}
	if typedish == nil {
		return nil
	}

	menudish := make(TypeDishListinMenu)
	dishlist := make(DichList, 0)

	for _, v := range typedish {
		for _, s := range dish {
			if v.ID == s.Meta.TypeDishID {
				dishlist = append(dishlist, NewDm(s))
			}
		}
		menudish[v.NameType] = dishlist
		dishlist = nil
	}
	return &menudish
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
