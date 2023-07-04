package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/internal"
)

var check = &cobra.Command{
	Use:   "check",
	Short: "check this program can execution",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var res = ""

		if internal.EnableCPUVirtualization() != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check CPU virtualization...   " + res + "\n")

		if internal.CanExecVirsh() != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check virsh command...        " + res + "\n")

		if internal.CanExecVirtInstall() != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check virt-install command... " + res + "\n")
	},
}
