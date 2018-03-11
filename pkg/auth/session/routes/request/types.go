package request

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"io"
	"io/ioutil"
	"strings"
)

type SessionCreate struct {
	Login    *string `json:"login,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (s *SessionCreate) DecodeAndValidate(reader io.Reader) *errors.Err {
	var (
		err error
	)

	log.Debug("Request: Session: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Session: decode and validate data for creating err: %s", err)
		return errors.New("user").Unknown(err)
	}

	if err = json.Unmarshal(body, s); err != nil {
		log.Errorf("Request: Session: convert struct from json err: %s", err)
		return errors.New("user").IncorrectJSON(err)
	}

	if s.Login == nil || *s.Login == "" {
		log.Errorf("Request: Session: parameter login not valid")
		return errors.New("user").BadParameter("login", err)
	}

	if s.Password == nil || *s.Password == "" {
		log.Errorf("Request: Session: parameter password not valid")
		return errors.New("user").BadParameter("password", err)
	}

	*s.Login = strings.ToLower(*s.Login)

	return nil
}
