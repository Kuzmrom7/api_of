package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
)

type RequestPersonCreate struct {
	Fio            string `json:"fio,omitempty"`
	Phone          string `json:"phone,omitempty"`
	NameTypePerson string `json:"nametypeperson"`
}

func (s *RequestPersonCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("personal").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("personal").IncorrectJSON(err)
	}

	if s.Fio == "" {
		return errors.New("personal").BadParameter("fio")
	}

	if s.Phone == "" {
		return errors.New("personal").BadParameter("phonenumber")
	}

	if s.NameTypePerson == "" {
		return errors.New("personal").BadParameter("nametypeperson")
	}

	return nil
}
