package types

import "time"

type Dish struct {
	Meta DishMeta `json:"meta"`
}

type DishMeta struct {
	ID         string    `json:"id"`
	Name       string    `json:"name,omitempty"`
	Desc       string    `json:"description,omitempty"`
	Timemin    int64     `json:"timemin,omitempty"`
	Url        string    `json:"url"`
	TypeDishID string    `json:"typedish"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

type TypeDishes struct {
	ID       string `json:"id"`
	NameType string `json:"nametype"`
}
