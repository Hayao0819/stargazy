package build

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "build",
		Short: "Build kernel",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	return &cmd
}
