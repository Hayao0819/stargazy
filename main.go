package main

import (
	"github.com/Hayao0819/stargazy/cmd"
	"github.com/Hayao0819/stargazy/utils"
)

func main() {
	root := cmd.Root()
	err := root.Execute()
	utils.ExitWithErr(root, err)
}
