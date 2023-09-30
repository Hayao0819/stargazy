package cmd

import "github.com/spf13/cobra"

func Root() *cobra.Command {
	return &cobra.Command{
		Use: "stargazy",
		Short: "Linux kernel manager",
		Long: "Stargazy is the tool to build, manage and backup kernels",
		
	}
}
