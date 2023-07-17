package kvm

import (
	"fmt"
	"os/exec"
)

// CheckVMStatus is checking vm name and status
func CheckVMStatus(name string) (string, error) {
	err := exec.Command("sh", "-c", "virsh list --all | grep "+name).Run()

	if err == nil {
		err := exec.Command("sh", "-c", "virsh list | grep "+name).Run()
		if err == nil {
			return "inactive", nil
		} else {
			return "active", nil
		}
	} else if err.Error() == "exit status 1" {
		return "NotFound", nil
	} else {
		return "", fmt.Errorf("command exection error for checking vm status")
	}
}
