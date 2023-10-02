package main

import (
	"os"

	"github.com/Hayao0819/stargazy/tools/cmd"
	"github.com/spf13/cobra"
)

func root() *cobra.Command {
	root := cobra.Command{
		Use:   "sg-tools",
		Short: "Developing helper",
	}

	root.AddCommand(cmd.GenConfigCmd(), cmd.RunCmd())

	return &root
}

func main() {
	root := root()
	if err := root.Execute(); err != nil {
		root.PrintErr(err)
		os.Exit(1)
	}
}
