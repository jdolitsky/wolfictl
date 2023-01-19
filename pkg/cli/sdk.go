package cli

import (
	"github.com/spf13/cobra"
)

func SDK() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "sdk",
		SilenceErrors: true,
		Short:         "Working with the Wolfi SDK",
	}

	cmd.AddCommand(SDKInit())
	return cmd
}
