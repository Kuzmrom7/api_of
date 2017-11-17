package types

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var(
	SecretAccessToken	= ""
)

//Generate new session structure
func NewSession(uid, username, email string) *Session{
	return &Session{
		Uid: uid,
		Username: username,
		Email: email,
	}
}

//Session data
type Session struct{
	Uid 			string
	Username 	string
	Email 		string
}


func (s *Session) Encode() (string, error) {
	context := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": s.Uid,
		"em":  s.Email,
		"un":  s.Username,
		"jti": time.Now().Add(time.Hour * 2232).Unix(),
		"exp": time.Now().Add(time.Hour * 2232).Unix(),
	})
	return context.SignedString([]byte(SecretAccessToken))
}
