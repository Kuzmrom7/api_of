package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewUser(obj *types.User) *User {
	return newUser(obj)
}
