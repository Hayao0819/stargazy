package lib

import (
	"os"
	"path"
)

var Pwd string = ""

func GetDebugConfPath() string {
	return path.Join(Pwd, "debug.json")
}

func init() {
	Pwd, _ = os.Getwd()
}
