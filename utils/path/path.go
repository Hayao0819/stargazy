package pathutils

import (
	"path/filepath"
	"strings"
)

func FromSlash(path ...string) string {

	return filepath.FromSlash(filepath.Clean(strings.Join(path, "/")))
}
