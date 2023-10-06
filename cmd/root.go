package cmd

import (
	"github.com/Hayao0819/stargazy/cmd/backup"
	"github.com/Hayao0819/stargazy/cmd/debug"
	"github.com/Hayao0819/stargazy/cmd/get"
	"github.com/Hayao0819/stargazy/cmd/list"
	"github.com/Hayao0819/stargazy/conf"
	errutils "github.com/Hayao0819/stargazy/utils/error"

	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "stargazy",
		Short:         "Linux kernel manager",
		Long:          "Stargazy is the tool to build, manage and backup kernels",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	var configFile string = ""
	cmd.PersistentFlags().StringVarP(&configFile, "config", "", "", "Specify stargazy config")
	cmd.AddCommand(backup.Root(), list.Root(), get.Root(), debug.Root())

	cobra.OnInitialize(func() {
		errutils.RunFuncWithErrorExit(cmd, func() error { return conf.Initilize(configFile) })
		errutils.RunFuncWithErrorExit(cmd, conf.Validate)
	})

	return cmd
}
