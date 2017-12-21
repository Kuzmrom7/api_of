package v1

import "time"

type Personal struct {
	Fio    string    `json:"name,omitempty"`
	Phone   string    `json:"phone,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type TypePersonalList []*TypePersonal

type TypePersonal struct {
	Meta TypePersonalMeta `json:"meta"`
}

type TypePersonalMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
