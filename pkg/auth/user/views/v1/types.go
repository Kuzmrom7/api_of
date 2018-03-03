package v1

type User struct {
	ID       string `json:"id"`
	Gravatar string `json:"gravatar"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
