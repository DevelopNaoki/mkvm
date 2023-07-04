package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mkvm",
	Short: "mkvm is manages and builds VMs with mkvm.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		return fmt.Errorf("need subcommand")
	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(
		list,
		version,
		check,
		build,
	)
	list.Flags().Bool("inactive", false, "list inactive domain")
}
