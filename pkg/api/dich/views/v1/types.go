package v1

type Dich struct {
	Id      string     `json:"id"`
	Name    string     `json:"name,omitempty"`
	Desc    string     `json:"description,omitempty"`
	Urls    []*UrlOpt  `json:"urls,omitempty"`
	Specs   []*SpecOpt `json:"specs,omitempty"`
	Timemin int64      `json:"timemin"`
}

type UrlOpt struct {
	Url string `json:"url"`
}

type SpecOpt struct {
	Size  string `json:"size"`
	Price string `json:"price"`
}

type DichList []*Dich

type TypeDishListinMenu map[string]DichList

type TypeDishList []*TypeDish

type TypeDish struct {
	Meta TypeDishMeta `json:"meta"`
}

type TypeDishMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
