package cmd

import (
	"github.com/Hayao0819/stargazy/cmd/backup"
	"github.com/Hayao0819/stargazy/conf"
	"github.com/Hayao0819/stargazy/utils"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stargazy",
		Short: "Linux kernel manager",
		Long:  "Stargazy is the tool to build, manage and backup kernels",
	}

	var configFile string = ""
	cmd.PersistentFlags().StringVarP(&configFile, "config", "", "", "Specify stargazy config")
	cmd.AddCommand(backup.Root())

	cobra.OnInitialize(func() {
		if _, err := conf.Initilize(configFile); err != nil {
			utils.ExitWithErr(cmd, err)
		}
	})

	return cmd
}
