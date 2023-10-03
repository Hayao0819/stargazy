package kernel

import (
	"fmt"

	"github.com/Hayao0819/stargazy/backup"
	errutils "github.com/Hayao0819/stargazy/utils/error"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var ignoreErr bool

	cmd := cobra.Command{
		Use:     "kernel",
		Args:    cobra.MaximumNArgs(1),
		Aliases: []string{"k"},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 && (!ignoreErr) {
				// Ask kernel path
				return fmt.Errorf("please specify a kernel path")
			}
			return nil
		},

		RunE: func(cmd *cobra.Command, args []string) error {
			list := backup.KernelBackupList{}
			errutils.RunFuncWithErrorExit(cmd, list.Initilize)

			return list.Add("hoge", args...)

		},
	}

	cmd.Flags().BoolVarP(&ignoreErr, "ignore-error", "", false, "Ignore argument error")
	// nolint:errcheck
	cmd.Flags().MarkHidden("ignore-error")
	return &cmd
}
