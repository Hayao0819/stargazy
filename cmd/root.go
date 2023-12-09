package cmd

import (
	"github.com/Hayao0819/stargazy/cmd/backup"
	"github.com/Hayao0819/stargazy/cmd/build"
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
	cmd.AddCommand(backup.Root(), list.Root(), get.Root(), debug.Root(), build.Root())

	// Config flags
	cmd.PersistentFlags().String("backup-dir", "", "Specify backup directory")
	cmd.PersistentFlags().String("kernel-source-dir", "", "Specify kernel source directory")

	errutils.RunFuncWithErrorExit(cmd, func() error {
		return conf.Viper.BindPFlag("backup_dir", cmd.PersistentFlags().Lookup("backup-dir"))
	})
	errutils.RunFuncWithErrorExit(cmd, func() error {
		return conf.Viper.BindPFlag("kernel_source_dir", cmd.PersistentFlags().Lookup("backup-dir"))
	})

	// Load config
	cobra.OnInitialize(func() {
		errutils.RunFuncWithErrorExit(cmd, func() error { return conf.Initilize(configFile) })
		errutils.RunFuncWithErrorExit(cmd, conf.Validate)
	})

	return cmd
}
