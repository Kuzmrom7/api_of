package generator

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"strings"
)

func Generatepassword(password string, salt string) (string, error) {
	pass := []byte(password + salt)

	hash, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func GenerateGravatar(email string) string {
	m := md5.New()
	if _, err := io.WriteString(m, strings.ToLower(email)); err != nil {
		return ""
	}
	return fmt.Sprintf("%x", m.Sum(nil))
}

func GenerateSalt(password string) (string, error) {
	buf := make([]byte, 10, 10+sha1.Size)
	_, err := io.ReadFull(rand.Reader, buf)
	if err != nil {
		fmt.Printf("random read fail: %v", err)
	}

	hash := sha1.New()
	_, err = hash.Write(buf)
	if err != nil {
		return "", err
	}

	_, err = hash.Write([]byte(password))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(buf)), nil
}
