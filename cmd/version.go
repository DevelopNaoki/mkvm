package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = &cobra.Command{
	Use:   "version",
	Short: "Print mkvm version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("mkvm version 1.1.0a\n")
	},
}
