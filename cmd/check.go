package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/DevelopNaoki/mkvm/internal"
)

var check = &cobra.Command{
	Use:   "check",
	Short: "check this program can execution",
	Args:  cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var res = ""

		err1 := internal.EnableCPUVirtualization()
		if err1 != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check CPU virtualization...   " + res + "\n")

		err2 := internal.CanExecVirsh()
		if err2 != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check virsh command...        " + res + "\n")

		err3 := internal.CanExecVirtInstall()
		if err3 != nil {
			res = internal.FAIL
		} else {
			res = internal.PASS
		}
		fmt.Print("Check virt-install command... " + res + "\n")

		if (err1 != nil) || (err2 != nil) || (err3 != nil) {
			os.Exit(-1)
		}
	},
}
