package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
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

	d.Specs = make([]*SpecOpt, 0)
	if obj.Specs != nil {
		for _, sp := range obj.Specs {
			s := new(SpecOpt)

			s.Size = sp.Size
			s.Price = sp.Price

			d.Specs = append(d.Specs, s)
		}
	}

	d.Timemin = obj.Meta.Timemin

	return d
}

func (obj *Dich) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.Dish) *Dich {
	d := new(Dich)

	d.Id = obj.Meta.ID
	d.Name = obj.Meta.Name
	d.Desc = obj.Meta.Desc
	d.TypeDishID = obj.Meta.TypeDishID
	d.Urls = make([]*UrlOpt, 0)
	if obj.Urls != nil {
		for _, url := range obj.Urls {
			u := new(UrlOpt)

			u.Url = url.Url

			d.Urls = append(d.Urls, u)
		}
	}

	if obj.Specs != nil {
		for _, sp := range obj.Specs {
			s := new(SpecOpt)

			s.Size = sp.Size
			s.Price = sp.Price

			d.Specs = append(d.Specs, s)
		}
	}
	d.Timemin = obj.Meta.Timemin

	return d
}

func NewDm(obj *types.Dish) *Dich {
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
	if obj.Specs != nil {
		for _, sp := range obj.Specs {
			s := new(SpecOpt)

			s.Size = sp.Size
			s.Price = sp.Price

			d.Specs = append(d.Specs, s)
		}
	}
	d.Timemin = obj.Meta.Timemin

	return d
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
