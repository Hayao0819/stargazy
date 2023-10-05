package cmd

import (
	"path"

	"github.com/Hayao0819/stargazy/tools/lib"
	executils "github.com/Hayao0819/stargazy/utils/exec"
	"github.com/spf13/cobra"
)

func RunCmd() *cobra.Command {
	withDebugConfig := false

	cmd := cobra.Command{
		Use:   "run",
		Args:  cobra.ArbitraryArgs,
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
			//cmd.PrintErr("Run ", goargs)

			return executils.Run("go", goargs...)
		},
	}

	// Disable help
	cmd.Flags().SetInterspersed(false)
	cmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	cmd.PersistentFlags().Lookup("help").Hidden = true

	cmd.Flags().BoolVarP(&withDebugConfig, "debug", "d", withDebugConfig, `Run with "--config DEBUG_CONFIG"`)

	return &cmd
}
