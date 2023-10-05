package kernel

import (
	"github.com/Hayao0819/stargazy/backup"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var ignoreErr bool

	cmd := cobra.Command{
		Use: "kernel",
		//Args:    cobra.MaximumNArgs(1),
		Aliases: []string{"k"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			/*
				if len(args) != 1 && (!ignoreErr) {
					// Ask kernel path
					return fmt.Errorf("please specify a kernel path")
				}*/
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			baklist := backup.List{}
			if err := baklist.Initilize("kernel-backup"); err != nil {
				return err
			}

			return nil

		},
	}

	cmd.Flags().BoolVarP(&ignoreErr, "ignore-error", "", false, "Ignore argument error")
	// nolint:errcheck
	cmd.Flags().MarkHidden("ignore-error")
	return &cmd
}
