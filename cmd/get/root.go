package get

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get kernel source",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.PrintErrln("Sorry that his function is not implemented...")
		},
	}

	cmd.AddCommand()
	return &cmd
}
