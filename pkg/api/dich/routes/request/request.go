package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
)

type RequestDichCreate struct {
	Name     string `json:"name,omitempty"`
	Desc     string `json:"description,omitempty"`
	Timemin  int64  `json:"timemin,omitempty"`
	TypeDish string `json:"typedish"`
	Url      string `json:"url"`
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

	if s.Name == "" {
		return errors.New("dish").BadParameter("name")
	}

	if s.Desc == "" {
		return errors.New("dish").BadParameter("name")
	}

	if s.TypeDish == "" {
		return errors.New("dish").BadParameter("typedish")
	}

	if s.Url == "" {
		return errors.New("dish").BadParameter("url")
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
	if s.Name == "" {
		return errors.New("dish").BadParameter("name")
	}

	return nil
}
