package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print kvm-compose version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("kvm-compose version 1.0.0a\n")
		return nil
	},
}
