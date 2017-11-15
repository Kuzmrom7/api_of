package pgsql


const (
	sqlGetUsers = `
		SELECT * FROM users
	`

	sqlCreqteUser = `
		INSERT INTO users (username, email) VALUES ($1, $2)
	`
)

type UserModel struct {

}

func GetUser(){

}