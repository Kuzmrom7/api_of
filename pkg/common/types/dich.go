package types

import "time"

type Dich struct {
	Meta DichMeta `json:"meta"`
}

type DichMeta struct {
	ID      string    `json:"id"`
	Name    string    `json:"name,omitempty"`
	Desc    string    `json:"description,omitempty"`
	Timemin int64     `json:"timemin,omitempty"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}
