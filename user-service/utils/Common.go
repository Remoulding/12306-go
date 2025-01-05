package utils

import (
	"math"
	"strings"
)

func IsEmail(str string) bool {
	return strings.Contains(str, "@")
}

func HashCode(str string) int {
	hash := 0
	for i := 0; i < len(str); i++ {
		hash = 31*hash + int(str[i])
	}
	return hash
}

func HashShardingIdx(username string) int {
	hash := HashCode(username)
	return int(math.Abs(float64(hash % UserRegisterReuseShardingCount)))
}
