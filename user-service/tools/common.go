package tools

import (
	"strings"
)

func IsEmail(str string) bool {
	return strings.Contains(str, "@")
}

func IsPhone(str string) bool {
	if len(str) != 11 {
		return false
	}
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}
