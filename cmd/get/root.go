package get

import (
	"github.com/Hayao0819/stargazy/kernel"
	"github.com/spf13/cobra"
)

func Root() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "Get kernel source from configurated upstreams",
		Long:  "Get kernel sources from various locations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			upstreams := kernel.GetUpstreamsFromConfig()
			for _, u := range upstreams {
				if u.Name == args[0] {
					if err := u.Get(); err != nil {
						return err
					}
				}
			}
			return nil
		},
	}

	cmd.AddCommand()
	return &cmd
}
