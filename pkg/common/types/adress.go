package types

type Adress struct {
	Meta AdressMeta `json:"meta"`
}

type AdressMeta struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	PlaceID  string `json:"place,omitempty"`
}
