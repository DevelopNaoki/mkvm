package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mkvm",
	Short: "mkvm is manages and builds VMs with mkvm.yml",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("error: need subcommand\n")
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
