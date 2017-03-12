package util

import (
	"regexp"
)

func MatchPhoneNum(phone string) bool {
	pattern := `1[3|5|7|8]\d{9}`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(phone)
}
