package main

import (
	"os"

	"github.com/spf13/cobra"
)

func root() *cobra.Command {
	return &cobra.Command{
		Use:   "sg-tools",
		Short: "Developing helper",
	}
}

func main() {
	root := root()
	if err := root.Execute(); err != nil {
		root.PrintErr(err)
		os.Exit(1)
	}
}
