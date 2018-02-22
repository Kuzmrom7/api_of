package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
)

type MenuCreate struct {
	Name      string `json:"name,omitempty"`
	Id_place 	string `json:"id_place"`
	Url       string `json:"url"`
}

type RequestMenuDishCreateRemove struct {
	NameMenu string `json:"namemenu,omitempty"`
	NameDish string `json:"namedish"`
}

type RequestMenuFetch struct {
	NameMenu string `json:"namemenu,omitempty"`
}

func (s *MenuCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

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

	if s.Url == "" {
		return errors.New("menu").BadParameter("url")
	}

	if s.Id_place == "" {
		return errors.New("menu").BadParameter("id_place")
	}

	return nil
}

func (s *RequestMenuDishCreateRemove) DecodeAndValidate(reader io.Reader) *errors.Err {

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

func (s *RequestMenuFetch) DecodeAndValidate(reader io.Reader) *errors.Err {

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

	if s.NameMenu == "" {
		return errors.New("menu").BadParameter("namemenu")
	}

	return nil
}

