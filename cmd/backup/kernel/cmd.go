package kernel

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	return &cobra.Command{
		Use:     "kernel",
		Args:    cobra.MaximumNArgs(1),
		Aliases: []string{"k"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				// Ask kernel path
				return fmt.Errorf("please specify a kernel path")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
