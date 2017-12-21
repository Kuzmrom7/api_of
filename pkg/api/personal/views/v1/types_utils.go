package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newPersonal(obj *types.Personal) *Personal{
	p := new(Personal)
	p.Fio = obj.Meta.Fio
	p.Phone = obj.Meta.Phone
	return p
}

func (obj *Personal) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}

func New(obj *types.TypePersonals) *TypePersonal {
	i := new(TypePersonal)
	i.Meta.Name = obj.NameType
	i.Meta.ID = obj.ID
	return i
}