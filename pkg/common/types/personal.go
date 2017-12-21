package types

import "time"

type Personal struct {
	Meta PersonalMeta `json:"meta"`
}

type PersonalMeta struct {
	ID             string    `json:"id"`
	Fio            string    `json:"fio,omitempty"`
	Phone          string    `json:"phone,omitempty"`
	TypePersonalID string    `json:"typeplace,omitempty"`
	PlaceID        string    `json:"user,omitempty"`
	Created        time.Time `json:"created"`
	Updated        time.Time `json:"updated"`
}

type TypePersonals struct {
	ID       string `json:"id"`
	NameType string `json:"nametype"`
}
