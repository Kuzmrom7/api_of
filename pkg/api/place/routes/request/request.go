package request

import (
	"io"

	"github.com/orderfood/api_of/pkg/common/errors"
	"io/ioutil"
	"encoding/json"

)

type RequestPlaceCreate struct {
	Name          string `json:"name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Url           string `json:"url,omitempty"`
	City          string `json:"city,omitempty"`
	Adress        string `json:"adress,omitempty"`
	NameTypePlace string `json:"nametypeplace,omitempty"`
}

func (s *RequestPlaceCreate) DecodeAndValidate(reader io.Reader) *errors.Err {

	var (
		err error
	)

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return errors.New("user").Unknown(err)
	}
	err = json.Unmarshal(body, s)
	if err != nil {
		return errors.New("user").IncorrectJSON(err)
	}


	return nil
}
