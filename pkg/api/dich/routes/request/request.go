package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"github.com/orderfood/api_of/pkg/log"
	"io/ioutil"
	"encoding/json"
)

type DishCreate struct {
	Name     string   `json:"name,omitempty"`
	Desc     string   `json:"description,omitempty"`
	Timemin  int64    `json:"timemin,omitempty"`
	IdTypeDish string   `json:"idtypedish"`
	Urls     []UrlOpt `json:"urls"`
}

type UrlOpt struct {
	Url string `json:"url"`
}

type DishUpdate struct {
	Id    	 string `json:"id,omitempty"`
	Desc    *string `json:"description,omitempty"`
	Timemin *int64  `json:"timemin,omitempty"`
}

func (s *DishCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Dish: decode and validate data for creating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Dish: decode and validate data for creating err: %s", err)
		return errors.New("dish").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Dish: convert struct from json err: %s", err)
		return errors.New("dish").IncorrectJSON(err)
	}

	if s.Name == "" {
		log.Error("Request: Dish: parameter name dish can not be empty")
		return errors.New("dish").BadParameter("name")
	}

	if s.Desc == "" {
		log.Error("Request: Dish: parameter desc dish can not be empty")
		return errors.New("dish").BadParameter("desc")
	}

	if s.IdTypeDish == "" {
		log.Error("Request: Dish: parameter id type dish can not be empty")
		return errors.New("dish").BadParameter("idtypedish")
	}

	if len(s.Urls) == 0 {
		log.Error("Request: Dish: parameter urls dish can not be empty")
		return errors.New("dish").BadParameter("urls")
	}

	return nil
}

func (s *DishUpdate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	log.Debug("Request: Dish: decode and validate data for updating")

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Errorf("Request: Dish: decode and validate data for updating err: %s", err)
		return errors.New("dish").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		log.Errorf("Request: Dish: convert struct from json err: %s", err)
		return errors.New("dish").IncorrectJSON(err)
	}

	if s.Id == "" {
		log.Error("Request: Dish: parameter id dish can not be empty")
		return errors.New("dish").BadParameter("name")
	}

	return nil
}
