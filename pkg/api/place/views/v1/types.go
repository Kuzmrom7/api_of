package v1

type Place struct {
	Id         string       `json:"id,omitempty"`
	Name       string       `json:"name,omitempty"`
	Phone      string       `json:"phone,omitempty"`
	Url        string       `json:"url,omitempty"`
	City       string       `json:"city,omitempty"`
	Adresses   []AdressOpt  `json:"adresses,omitempty"`
	TypesPlace []TypePlaces `json:"typesplace"`
}

type AdressOpt struct {
	Adress string `json:"adress"`
}

type TypePlaces struct {
	ID       string `json:"id"`
	NameType string `json:"nametype"`
}

type TypePlaceList []*TypePlace

type TypePlace struct {
	Meta TypePlaceMeta `json:"meta"`
}

type TypePlaceMeta struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
