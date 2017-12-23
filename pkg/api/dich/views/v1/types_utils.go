package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newDich(obj *types.Dish) *Dich {
	d := new(Dich)
	d.Name = obj.Meta.Name
	d.Desc = obj.Meta.Desc
	d.Url = obj.Meta.Url
	d.Timemin = obj.Meta.Timemin
	d.Created = obj.Meta.Created
	d.Updated = obj.Meta.Updated

	return d
}

func (obj *Dich) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.Dish) *Dich {
	i := new(Dich)
	i.Name = obj.Meta.Name
	i.Desc = obj.Meta.Desc
	i.Url = obj.Meta.Url
	i.Timemin = obj.Meta.Timemin
	i.Updated = obj.Meta.Updated
	i.Created = obj.Meta.Created
	return i
}

func Newt(obj *types.TypeDishes) *TypeDish {
	i := new(TypeDish)
	i.Meta.Name = obj.NameType
	i.Meta.ID = obj.ID
	return i
}

func (obj *DichList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &DichList{}
	}
	return json.Marshal(obj)
}

func (obj *TypeDishList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &TypeDishList{}
	}
	return json.Marshal(obj)
}