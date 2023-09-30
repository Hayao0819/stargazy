package backup

import (
	"github.com/Hayao0819/stargazy/cmd/backup/kernel"
	"github.com/Hayao0819/stargazy/cmd/backup/source"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "backup",
		Short:   "Manage kernel backup",
		Aliases: []string{"bak", "b"},
	}

	cmd.AddCommand(source.Cmd(), kernel.Cmd())

	return cmd
}
