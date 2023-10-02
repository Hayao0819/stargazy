package cmd

import (
	"os/exec"
	"path"

	"github.com/Hayao0819/stargazy/tools/lib"
	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	withDebugConfig := false

	cmd := cobra.Command{
		Use:   "run",
		Args:  cobra.NoArgs,
		Short: "Run stargazy",
		RunE: func(cmd *cobra.Command, args []string) error {
			sgargs := []string{}
			goargs := []string{"run"}

			if withDebugConfig {
				json := lib.GetDebugConfPath()
				sgargs = append(sgargs, "--config", json)
			}

			goargs = append(goargs, path.Join(lib.Pwd, "main.go"))
			goargs = append(goargs, sgargs...)
			goargs = append(goargs, args...)

			return exec.Command("go", goargs...).Run()
		},
	}

	// Disable help
	cmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	cmd.PersistentFlags().Lookup("help").Hidden = true

	cmd.Flags().BoolVarP(&withDebugConfig, "debug", "d", withDebugConfig, `Run with "--config DEBUG_CONFIG"`)

	return &cmd
}
