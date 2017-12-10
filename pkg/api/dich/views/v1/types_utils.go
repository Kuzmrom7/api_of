package v1

import (
	"github.com/orderfood/api_of/pkg/common/types"
	"encoding/json"
)

func newDich(obj *types.Dich) *Dich {
	d := new(Dich)
	d.Name = obj.Meta.Name
	d.Desc = obj.Meta.Desc
	d.Created = obj.Meta.Created
	d.Updated = obj.Meta.Updated

	return d
}

func (obj *Dich) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}
