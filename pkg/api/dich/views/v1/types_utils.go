package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newDich(obj *types.Dish) *Dich {
	d := new(Dich)
	d.Id = obj.Meta.ID
	d.Name = obj.Meta.Name
	d.Desc = obj.Meta.Desc

	d.Urls = make([]*UrlOpt, 0)
	if obj.Urls != nil {
		for _, url := range obj.Urls {
			u := new(UrlOpt)

			u.Url = url.Url

			d.Urls = append(d.Urls, u)
		}
	}

	d.Timemin = obj.Meta.Timemin

	return d
}

func (obj *Dich) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.Dish) *Dich {
	i := new(Dich)

	i.Id = obj.Meta.ID
	i.Name = obj.Meta.Name
	i.Desc = obj.Meta.Desc
	i.Urls = make([]*UrlOpt, 0)
	if obj.Urls != nil {
		for _, url := range obj.Urls {
			u := new(UrlOpt)

			u.Url = url.Url

			i.Urls = append(i.Urls, u)
		}
	}
	i.Timemin = obj.Meta.Timemin

	return i
}

func NewDm(obj *types.Dish) *Dich {
	i := new(Dich)

	i.Id = obj.Meta.ID
	i.Name = obj.Meta.Name
	i.Desc = obj.Meta.Desc
	i.Urls = make([]*UrlOpt, 0)
	if obj.Urls != nil {
		for _, url := range obj.Urls {
			u := new(UrlOpt)

			u.Url = url.Url

			i.Urls = append(i.Urls, u)
		}
	}
	i.Timemin = obj.Meta.Timemin

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

func (obj *TypeDishListinMenu) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &TypeDishListinMenu{}
	}
	return json.Marshal(obj)
}
