package util

import "strings"

func IsEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) == 0
}

func IsNotEmptyString(input string) bool {
	return len(strings.TrimSpace(input)) > 0
}
