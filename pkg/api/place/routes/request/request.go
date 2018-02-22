package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"
	"github.com/orderfood/api_of/pkg/log"
)

type RequestPlaceCreate struct {
	Name          string `json:"name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Url           string `json:"url,omitempty"`
	City          string `json:"city,omitempty"`
	Adress        string `json:"adress,omitempty"`
	NameTypePlace string `json:"nametypeplace,omitempty"`
}

type RequestPlaceUpdate struct {
	Phone  *string `json:"phone,omitempty"`
	Url    *string `json:"url,omitempty"`
	City   *string `json:"city,omitempty"`
	Adress *string `json:"adress,omitempty"`
}

func (s *RequestPlaceCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

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

	if s.NameTypePlace == "" {
		log.Error("Request: Place: parameter type place can not be empty")
		return errors.New("place").BadParameter("nametypeplace")
	}

	return nil
}

func (s *RequestPlaceUpdate) DecodeAndValidate(reader io.Reader) *errors.Err {

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

	return nil
}
