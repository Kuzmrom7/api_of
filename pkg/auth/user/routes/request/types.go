package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
	"github.com/orderfood/api_of/pkg/util/validator"
	"strings"
	"github.com/orderfood/api_of/pkg/log"
)

type RequestUserCreate struct {
	Email    *string `json:"email,omitempty"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	//It is a struct for body data for account create route
	//Pointer is for data validating
}

func (s *RequestUserCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: User: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: User: decode and validate data for creating err: %s", err)
		return errors.New("user").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: User: convert struct from json err: %s", err)
		return errors.New("user").IncorrectJSON(err)
	}

	if s.Username == nil {
		log.Error("Request: User: parameter name can not be empty")
		return errors.New("user").BadParameter("username")
	}
	if !validator.IsUsername(*s.Username) {
		log.Error("Request: User: parameter name not valid")
		return errors.New("user").BadParameter("username")
	}

	if s.Email == nil {
		log.Error("Request: User: parameter email can not be empty")
		return errors.New("user").BadParameter("email")
	}

	if !validator.IsEmail(*s.Email) {
		log.Error("Request: User: parameter email not valid")
		return errors.New("user").BadParameter("email")
	}

	if s.Password == nil {
		log.Error("Request: User: parameter password can not be empty")
		return errors.New("user").BadParameter("password")
	}

	if !validator.IsPassword(*s.Password) {
		log.Error("Request: User: parameter password not valid")
		return errors.New("user").BadParameter("password")
	}

	*s.Username = strings.ToLower(*s.Username)
	*s.Email = strings.ToLower(*s.Email)

	return nil
}
