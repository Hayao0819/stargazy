package main

import (
	"github.com/Hayao0819/stargazy/cmd"
	errutils "github.com/Hayao0819/stargazy/utils/error"
)

func main() {
	root := cmd.Root()
	err := root.Execute()
	errutils.ExitWithErr(root, err)
}
