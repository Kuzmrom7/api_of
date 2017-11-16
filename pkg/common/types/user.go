package types

type User struct {
	Meta			 	UserMeta 				`json:"meta"`
	//Profile UserProfile `json:"profile"`
	Security 		UserSecurity  	`json:"security"`
}

type UserMeta struct {
	ID 					string `json:"id"`
	Username 		string `json:"username"`
	Email 			string `json:"email"`
	Gravatar 		string `json:"gravatar"`
	Active 			bool 	 `json:"active"`
}

type UserSecurity struct {
	Pass 				UserPassword `json:"pass"`
	SSH  				[]UserSSH    `json:"ssh"`
}

type UserPassword struct {
	Password 		string `json:"password"`
	Salt 				string `json:"salt"`
}

type UserSSH struct {
	Name        string `json:"name"`
	Fingerprint string `json:"fingerprint"`
	Key         string `json:"key"`
}