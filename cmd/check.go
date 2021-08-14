package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/process"
)

var check = &cobra.Command{
	Use:   "check",
	Short: "Check thin program can execution",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		process.Check()
		return nil
	},
}
