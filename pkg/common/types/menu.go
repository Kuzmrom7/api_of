package types

import "time"

type Menu struct {
	Meta MenuMeta `json:"meta"`
}

type MenuMeta struct {
	ID      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	PlaceID string    `json:"place,omitempty"`
	Url     string    `json:"url"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
