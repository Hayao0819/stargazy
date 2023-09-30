package main

import (
	"os"

	"github.com/Hayao0819/stargazy/cmd"
)

func main() {
	root := cmd.Root()
	if err := root.Execute(); err != nil {
		root.PrintErr(err)
		os.Exit(1)
	}
}
