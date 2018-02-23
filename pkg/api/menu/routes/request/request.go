package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"io/ioutil"
	"encoding/json"
)

type MenuCreate struct {
	Name     string `json:"name,omitempty"`
	Id_place string `json:"id_place"`
	Url      string `json:"url"`
}

func (s *MenuCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Menu: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Menu: decode and validate data for creating err: %s", err)
		return errors.New("menu").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Menu: convert struct from json err: %s", err)
		return errors.New("menu").IncorrectJSON(err)
	}

	if s.Name == "" {
		log.Error("Request: Menu: parameter name menu can not be empty")
		return errors.New("menu").BadParameter("name")
	}

	if s.Url == "" {
		log.Error("Request: Menu: parameter url can not be empty")
		return errors.New("menu").BadParameter("url")
	}

	if s.Id_place == "" {
		log.Error("Request: Menu: parameter id_place can not be empty")
		return errors.New("menu").BadParameter("id_place")
	}

	return nil
}
