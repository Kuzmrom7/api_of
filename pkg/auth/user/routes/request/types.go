package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
	"github.com/orderfood/api_of/pkg/util/validator"
	"strings"
)

type RequestUserCreate struct {
	Email *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	//It is a struct for body data for account create route
	//Pointer is for data validating
}

func (s *RequestUserCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var(
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("user").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return  errors.New("user").IncorrectJSON(err)
	}

	if s.Email == nil || !validator.IsEmail(*s.Email) {
		return errors.New("user").BadParameter("email")
	}
	if s.Username == nil || !validator.IsUsername(*s.Username) {
		return errors.New("user").BadParameter("username")
	}
	if s.Password == nil || !validator.IsPassword(*s.Password) {
		return errors.New("user").BadParameter("password")
	}

	*s.Username = strings.ToLower(*s.Username)
	*s.Email = strings.ToLower(*s.Email)

	return nil
}