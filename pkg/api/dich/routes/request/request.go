package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
)

type RequestDichCreate struct {
	Name    string `json:"name,omitempty"`
	Desc    string `json:"description,omitempty"`
	Timemin int64  `json:"timemin,omitempty"`
}

type RequestDichRemove struct {
	Name string `json:"name,omitempty"`
}

func (s *RequestDichCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("dich").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("dich").IncorrectJSON(err)
	}

	return nil
}

func (s *RequestDichRemove) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("dich").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("dich").IncorrectJSON(err)
	}

	return nil
}
