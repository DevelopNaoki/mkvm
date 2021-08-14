package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/process"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "build VM by mkvm.yml",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.Build()
		return nil
	},
}
