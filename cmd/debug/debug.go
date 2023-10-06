package debug

import (
	"github.com/Hayao0819/stargazy/backup"
	"github.com/spf13/cobra"
)

func BackupCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "backup",
		Args:  cobra.MinimumNArgs(1),
		Short: "Test backup package",
		Long:  "Backup any files and show data struct",
		RunE: func(cmd *cobra.Command, args []string) error {
			bak := backup.List{}
			if err := bak.Initilize("debug-backup"); err != nil {
				return err
			}
			if err := bak.Add(args...); err != nil {
				return err
			}

			cmd.Println(bak)
			for _, b := range bak.List {
				cmd.Println(*b)
			}
			return nil
		},
	}

	return &cmd
}

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:    "debug",
		Short:  "Subcommands for debug",
		Args:   cobra.NoArgs,
		Hidden: true,
	}
	cmd.AddCommand(BackupCmd())
	return &cmd
}
