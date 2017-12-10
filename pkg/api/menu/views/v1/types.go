package v1

import "time"

type Menu struct {
	Name      string    `json:"name,omitempty"`
	NamePlace string    `json:"nameplace,omitempty"`
	Created   time.Time `json:"created"`
	Updated   time.Time `json:"updated"`
}
