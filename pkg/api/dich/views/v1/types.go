package v1

import "time"

type Dich struct {
	Name    string    `json:"name,omitempty"`
	Desc    string    `json:"description,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type DichList []*Dich
