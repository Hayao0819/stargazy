package cmd

import (
	"os"
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

			// testdir
			testdir := path.Join(lib.Pwd, "testdir")

			// Set dirs
			backupDir := path.Join(testdir, "backup")
			kernelSourceDir := path.Join(testdir, "kernel_source")

			for _, dir := range []string{backupDir, kernelSourceDir} {
				if err := os.MkdirAll(dir, 0750); err != nil {
					return err
				}
			}

			// Set files
			viper.Set("backup_dir", backupDir)
			viper.Set("kernel_source_dir", kernelSourceDir)

			// Write
			return viper.WriteConfigAs(json)

			//return nil
		},
	}
	return &cmd
}
