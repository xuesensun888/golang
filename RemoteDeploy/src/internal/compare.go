package internal

import (
	"strconv"
	"strings"
)

func CompareVersion(compareFunc func(a, b string) bool, a, b string) bool {
	return compareFunc(a, b)
}

func IsNew(a, b string) bool {
	segement := 3
	newArray := strings.Split(a, ".")
	curArray := strings.Split(b, ".")
	if len(newArray) != segement || len(curArray) != segement {
		return false
	}

	for i := 0; i < segement; i++ {
		a, _ := strconv.ParseInt(newArray[i], 10, 32)
		b, _ := strconv.ParseInt(curArray[i], 10, 32)
		if a > b {
			return true
		}

	}
	return false
}
