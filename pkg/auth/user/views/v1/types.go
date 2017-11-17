package v1

//import "time"

type User struct {
	ID 					string 		`json:"id"`
	Gravatar 		string 		`json:"gravatar"`
	Email				string 		`json:"email"`
	Username		string		`json:"username"`
	//Created			time.Time `json:"created"`
	//Updated			time.Time `json:"updated"`
}
