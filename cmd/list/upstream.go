package list

import (
	"github.com/Hayao0819/stargazy/kernel"
	"github.com/spf13/cobra"
)

func UpstreamList() *cobra.Command{
	cmd := cobra.Command{
		Use: "upstream",
		Short: "List upstream",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			for _,u := range kernel.GetUpstreamsFromConfig(){
				cmd.Println(u.Name)
			}
			return nil
		},
	}
	return &cmd
}
