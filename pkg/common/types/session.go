package types

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SecretAccessToken         = ""
	ErrUnexpectedSigninMethod = errors.New("UNEXPECTED_SIGNIN_METHD")
	ErrSessionTokenHasNoEXP   = errors.New("NO_EXP_IN_TOKEN")
	ErrSessionTokenHasNoJTI   = errors.New("NO_JTI_IN_TOKEN")
)

//Generate new session structure
func NewSession(uid, username, email string) *Session {
	return &Session{
		Uid:      uid,
		Username: username,
		Email:    email,
	}
}

//Session data
type Session struct {
	Uid      string
	Username string
	Email    string
}

func (s *Session) Decode(token string) error {

	payload, err := jwt.Parse(token, func(payload *jwt.Token) (interface{}, error) {
		result := []byte(SecretAccessToken)

		err := func(token *jwt.Token) error {

			claims := token.Claims.(jwt.MapClaims)

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return ErrUnexpectedSigninMethod
			}

			if claims["exp"] == nil {
				return ErrSessionTokenHasNoEXP
			}

			if claims["jti"] == nil {
				return ErrSessionTokenHasNoJTI
			}

			return nil
		}(payload)

		return result, err
	})

	if err != nil || !payload.Valid {
		return err
	}

	claims := payload.Claims.(jwt.MapClaims)

	if _, ok := claims["uid"]; ok {
		s.Uid = claims["uid"].(string)
	}
	if _, ok := claims["em"]; ok {
		s.Email = claims["em"].(string)
	}
	if _, ok := claims["un"]; ok {
		s.Username = claims["un"].(string)
	}

	return nil
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
