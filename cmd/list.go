package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/process"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "display VM status",
	Args:  cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			var option string
			if len(args) == 0 || args[0] == "active" {
				option = "active"
			} else if args[0] == "all" || args[0] == "inactive" {
				option = args[0]
			} else {
				return nil
			}
			process.VMList(option)
		} else {
			return nil
		}
		return nil
	},
}
