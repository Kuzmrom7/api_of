package v1

import "github.com/orderfood/api_of/pkg/common/types"

func NewSession(obj *types.Session) *Session {
	return New(obj)
}
