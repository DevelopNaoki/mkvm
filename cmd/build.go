package cmd

import (
	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/process"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "build VM with mkvm.yml",
	RunE: func(cmd *cobra.Command, args []string) error {
		process.Build()
		return nil
	},
}
