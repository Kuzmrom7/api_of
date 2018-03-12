package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
)

func newPersonal(obj *types.Personal) *Personal {
	p := new(Personal)
	p.Fio = obj.Meta.Fio
	p.Phone = obj.Meta.Phone
	return p
}

func Newt(obj *types.TypePersonals) *TypePersonal {
	i := new(TypePersonal)
	i.Meta.Name = obj.NameType
	i.Meta.ID = obj.ID
	return i
}

func New(obj *types.Personal) *Personal {
	i := new(Personal)
	i.Fio = obj.Meta.Fio
	i.Phone = obj.Meta.Phone
	i.Updated = obj.Meta.Updated
	i.Created = obj.Meta.Created
	return i
}

func (obj *Personal) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func (obj *TypePersonalList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &TypePersonalList{}
	}
	return json.Marshal(obj)
}

func (obj *PersonalList) ToJson() ([]byte, error) {
	if obj == nil {
		obj = &PersonalList{}
	}
	return json.Marshal(obj)
}
