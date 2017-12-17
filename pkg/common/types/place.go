package types

import "time"

type Place struct {
	Meta PlaceMeta `json:"meta"`
	//Type TypePlaces `json:"type"`
}

type PlaceMeta struct {
	ID          string    `json:"id"`
	Name        string    `json:"name,omitempty"`
	Phone       string    `json:"phone,omitempty"`
	Url         string    `json:"url,omitempty"`
	City        string    `json:"city,omitempty"`
	Adress      string    `json:"adress,omitempty"`
	TypePlaceID string    `json:"typeplace,omitempty"`
	UserID      string    `json:"user,omitempty"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
}

type TypePlaces struct {
	ID       string `json:"id"`
	NameType string `json:"nametype"`
}
