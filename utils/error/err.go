package errutils

import (
	"os"

	"github.com/spf13/cobra"
)

func ExitWithErr(cmd *cobra.Command, err error) {
	if err == nil {
		os.Exit(0)
	} else {
		cmd.PrintErrln(err)
		os.Exit(1)
	}
}

func RunFuncWithErrorExit(cmd *cobra.Command, fn func() error) {
	err := fn()
	if err != nil {
		ExitWithErr(cmd, err)
	}
}
