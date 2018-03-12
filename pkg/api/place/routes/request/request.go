package request

import (
	"io"

	"encoding/json"
	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"io/ioutil"
)

type PlaceCreate struct {
	Name       string         `json:"name,omitempty"`
	TypesPlace []TypePlaceOpt `json:"typesplace,omitempty"`
}

type TypePlaceOpt struct {
	IdTypePlace   string `json:"idtypeplace"`
	NameTypePlace string `json:"nametypeplace"`
}

type PlaceUpdate struct {
	Id    string  `json:"id,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Url   *string `json:"url,omitempty"`
	City  *string `json:"city,omitempty"`
}

func (s *PlaceCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Place: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Place: decode and validate data for creating err: %s", err)
		return errors.New("place").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Place: convert struct from json err: %s", err)
		return errors.New("place").IncorrectJSON(err)
	}

	if s.Name == "" {
		log.Error("Request: Place: parameter name place can not be empty")
		return errors.New("place").BadParameter("name")
	}

	if len(s.TypesPlace) == 0 {
		log.Error("Request: Place: parameter type place can not be empty")
		return errors.New("place").BadParameter("typepslace")
	}

	return nil
}

func (s *PlaceUpdate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Place: decode and validate data for updating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Place: decode and validate data for updating err: %s", err)
		return errors.New("place").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Place: convert struct from json err: %s", err)
		return errors.New("place").IncorrectJSON(err)
	}

	if s.Id == "" {
		log.Error("Request: Place: parameter id place can not be empty")
		return errors.New("place").BadParameter("name")
	}

	return nil
}
