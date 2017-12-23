package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
)

type RequestMenuCreate struct {
	Name      string `json:"name,omitempty"`
	NamePlace string `json:"nameplace"`
	Url       string `json:"url"`
	//	Name          string `json:"name,omitempty"`
}

type RequestMenuDishCreate struct {
	NameMenu string `json:"namemenu,omitempty"`
	NameDish string `json:"namedish"`
}

func (s *RequestMenuCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("menu").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("menu").IncorrectJSON(err)
	}

	if s.Name == "" {
		return errors.New("menu").BadParameter("name")
	}

	if s.NamePlace == "" {
		return errors.New("menu").BadParameter("nameplace")
	}
	if s.Url == "" {
		return errors.New("menu").BadParameter("url")
	}

	return nil
}


func (s *RequestMenuDishCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("menu").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("menu").IncorrectJSON(err)
	}

	if s.NameDish == "" {
		return errors.New("menu").BadParameter("namedish")
	}

	if s.NameMenu == "" {
		return errors.New("menu").BadParameter("namemenu")
	}

	return nil
}
