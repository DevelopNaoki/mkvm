package internal

import (
	"fmt"
	"os/exec"
)

var PASS = "[\x1b[32mPass\x1b[0m]"
var FAIL = "[\x1b[31mFail\x1b[0m]"

// EnableCPUVirtualization checks whether cpu virtualization is enabled.
func EnableCPUVirtualization() error {
	cmd := "grep -E 'svm|vmx' /proc/cpuinfo"
	err := exec.Command("sh", "-c", cmd).Run()
	if err != nil {
		return fmt.Errorf("failed get cpu virtualization or disable cpu virtualization")
	}
	return nil
}

// CanExecVirsh checks if virsh command is executable
func CanExecVirsh() error {
	err := exec.Command("virsh", "version").Run()
	if err != nil {
		return fmt.Errorf("failed execute virsh command")
	}
	return nil
}

// CanExecVirtInstall checks if virt-install command is executable
func CanExecVirtInstall() error {
	err := exec.Command("virt-install", "--version").Run()
	if err != nil {
		return fmt.Errorf("failed execute virsh-install command")
	}
	return nil
}

// CheckFile is checking file exsist
func CheckFile(path string) (string, error) {
	err := exec.Command("find", path).Run()

	if err != nil {
		return "exist", nil
	} else if err.Error() == "exit code 1" {
		return "doesNotExist", nil
	} else {
		return "", fmt.Errorf("command exection error for checking vm status")
	}
}
