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

// 大文字と小文字を無視して文字列を比較します
func Match(pattern, name string) bool {
	return strings.ToLower(pattern) == strings.ToLower(name)
}
