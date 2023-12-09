package list

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "SHow the various list",
	}

	cmd.AddCommand(BackUpList(), UpstreamList())

	return cmd
}
