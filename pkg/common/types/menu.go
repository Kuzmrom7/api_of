package types


type Menu struct {
	Meta MenuMeta  `json:"meta"`
}
type MenuMeta struct {
	ID_menu		string `json:"_menu"`
	Name_menu	string  `json:"name_menu"`
	Created 	string  `json:"created"`
	Updated 	string  `json:"updated"`
}

type MenuAdd struct {
	ID_menu		string `json:"_menu"`
	Name_menu	string  `json:"name_menu"`
	Id_Place	string 	`json:"id_place"`
	User_id		string  `json:"user_id"`
	}