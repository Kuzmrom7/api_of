package request

import (
	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"io"
	"io/ioutil"
)

type AdressCreate struct {
	Name     string `json:"name,omitempty"`
	Id_place string `json:"id_place"`
}

func (s *AdressCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Adress: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Adress: decode and validate data for creating err: %s", err)
		return errors.New("menu").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Adress: convert struct from json err: %s", err)
		return errors.New("menu").IncorrectJSON(err)
	}

	if s.Name == "" {
		log.Error("Request: Adress: parameter name adress can not be empty")
		return errors.New("menu").BadParameter("name")
	}

	if s.Id_place == "" {
		log.Error("Request: Adress: parameter id_place can not be empty")
		return errors.New("menu").BadParameter("id_place")
	}

	return nil
}
