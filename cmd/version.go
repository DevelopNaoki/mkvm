package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print mkvm version",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Print("mkvm version 1.0.0a\n")
		return nil
	},
}
