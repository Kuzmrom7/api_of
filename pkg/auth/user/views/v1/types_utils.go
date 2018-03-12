package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
)

func newUser(obj *types.User) *User {
	u := new(User)
	u.ID = obj.Meta.ID
	u.Username = obj.Meta.Username
	u.Gravatar = obj.Meta.Gravatar
	u.Email = obj.Meta.Email
	return u
}

func (obj *User) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}
