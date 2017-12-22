package v1

import "time"

type Dich struct {
	Name    string    `json:"name,omitempty"`
	Desc    string    `json:"description,omitempty"`
	Url     string 		`json:"url"`
	Timemin int64 		`json:"timemin"`
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
}

type DichList []*Dich

type TypeDishList []*TypeDish

type TypeDish struct {
	Meta TypeDishMeta `json:"meta"`
}

type TypeDishMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}