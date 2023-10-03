package cmd

import (
	"path"

	"github.com/Hayao0819/stargazy/conf"
	"github.com/Hayao0819/stargazy/tools/lib"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GenConfigCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "genconfig",
		Short: "Generate stargazy config for debug",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Generate config
			if err := conf.Initilize(""); err != nil {
				return err
			}

			// Get debug.json
			json := lib.GetDebugConfPath()

			// Set backup_dir
			viper.Set("backup_dir", path.Join(lib.Pwd, "testdir", "backup"))

			// Write
			viper.WriteConfigAs(json)

			return nil
		},
	}
	return &cmd
}
