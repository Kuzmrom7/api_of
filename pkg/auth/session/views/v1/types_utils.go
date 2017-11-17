package v1

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/types"
)

func New(obj *types.Session) *Session{
	var token, err = obj.Encode()
	if err != nil {
		return nil
	}
	return &Session{token}
}

func (obj *Session) ToJson() ([]byte, error) {
	return json.Marshal(obj)
}