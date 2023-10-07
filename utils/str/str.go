package strutils

import (
	"strings"
)

func Join(elems ...string) string {
	return strings.Join(elems, "")
}

func IsEmpty(str string) bool {
	return strings.TrimSpace(str) == ""
}
