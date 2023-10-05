package list

import (
	"github.com/Hayao0819/stargazy/backup"
	"github.com/spf13/cobra"
)

func BackUpList() *cobra.Command {
	cmd := cobra.Command{
		Use:     "kernel-backup",
		Short:   "カーネルのバックアップ一覧",
		Aliases: []string{"kb"},
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			bak, err := backup.GetKernelBackupList()
			if err != nil {
				return err
			}
			for _, b := range bak.List {
				cmd.Println(b.Date.String())
				for _, f := range b.Files {
					cmd.Print("  ")
					cmd.Println(f.Original)
				}
			}
			return nil
		},
	}
	return &cmd
}
