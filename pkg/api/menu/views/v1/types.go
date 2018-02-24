package v1


type Menu struct {
	Id 				string 		`json:"id"`
	Name      string    `json:"name,omitempty"`
	Url 			string    `json:"url"`
}

type MenuList []*Menu