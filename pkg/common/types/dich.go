package types

import "time"

type Dish struct {
	Meta  DishMeta  `json:"meta"`
	Urls  []UrlsOpt `json:"urls,omitempty"`
	Specs []SpecOpt `json:"specs,omitempty"`
}

type UrlsOpt struct {
	Url string `json:"url"`
}

type SpecOpt struct {
	Size  string `json:"size"`
	Price string `json:"price"`
}

type DishMeta struct {
	ID         string    `json:"id"`
	Name       string    `json:"name,omitempty"`
	Desc       string    `json:"description,omitempty"`
	Timemin    int64     `json:"timemin,omitempty"`
	UserID     string    `json:"userid"`
	TypeDishID string    `json:"typedish"`
	Created    time.Time `json:"created"`
	Updated    time.Time `json:"updated"`
}

type TypeDishes struct {
	ID       string `json:"id"`
	NameType string `json:"nametype"`
}
