package request

import (
	"github.com/orderfood/api_of/pkg/common/errors"
	"io"
	"io/ioutil"
	"encoding/json"
	"strings"
)
type RequestSessionCreate struct {
	Login 		*string `json:"login,omitempty"`
	Password 	*string `json:"password,omitempty"`
}

func (s *RequestSessionCreate) DecodeAndValidate(reader io.Reader) *errors.Err{
	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("user").Unknown(err)
	}

	if err = json.Unmarshal(body, s); err != nil {
		return errors.New("user").IncorrectJSON(err)
	}

	if s.Login == nil || *s.Login == "" {
		return errors.New("user").BadParameter("login", err)
	}

	if s.Password == nil || *s.Password == "" {
		return errors.New("user").BadParameter("password", err)
	}

	*s.Login = strings.ToLower(*s.Login)

	return  nil
}