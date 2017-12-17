package v1

type Place struct {
	Name          string `json:"name,omitempty"`
	Phone         string `json:"phone,omitempty"`
	Url           string `json:"url,omitempty"`
	City          string `json:"city,omitempty"`
	Adress        string `json:"adress,omitempty"`
	NameTypePlace string `json:"nametypeplace,omitempty"`
}

type TypePlace struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type TypePlaceList []*TypePlace
