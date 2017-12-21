package validator

import (
	"github.com/asaskevich/govalidator"
	"regexp"
)

func IsEmail(s string) bool {
	return govalidator.IsEmail(s)
}

func IsUsername(s string) bool {
	reg, _ := regexp.Compile("[A-Za-z0-9]+(?:[_-][A-Za-z0-9]+)*")
	str := reg.FindStringSubmatch(s)
	if len(str) == 1 && str[0] == s && len(s) >= 4 && len(s) <= 64 {
		return true
	}
	return false
}

func IsPassword(s string) bool {
	return len(s) > 6
}
