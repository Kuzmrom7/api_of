package v1

type Adress struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	PlaceID  string `json:"place"`
	MemberID string `json:"member"`
}

type AdressList []*Adress
