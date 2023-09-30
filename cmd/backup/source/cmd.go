package source

import "github.com/spf13/cobra"

func Cmd()*cobra.Command{
	return &cobra.Command{
		Use: "source",
	}
}
