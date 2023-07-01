package internal

import (
	"fmt"
	"os"
	"os/exec"
)

// Check is the caller of each checking function
func Check() error {
	var status string
	TRUE := "[\x1b[32mPass\x1b[0m]"
	FALSE := "[\x1b[31mFalse\x1b[0m]"

	err1 := EnableCPUVirtualization()
	if err1 != nil {
		status = FALSE
	} else {
		status = TRUE
	}
	fmt.Print("Check CPU virtualization...   " + status + "\n")

	err2 := CanExecVirsh()
	if err2 != nil {
		status = FALSE
	} else {
		status = TRUE
	}
	fmt.Print("Check virsh command...        " + status + "\n")

	err3 := CanExecVirtInstall()
	if err3 != nil {
		status = FALSE
	} else {
		status = TRUE
	}
	fmt.Print("Check virt-install command... " + status + "\n")

	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("failed pass check")
	}
	return nil
}

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

// CheckVMStatus is checking vm name and status
func CheckVMStatus(name string) (status string) {
	err := exec.Command("sh", "-c", "virsh list --all | grep "+name).Run()

	if err == nil {
		err := exec.Command("sh", "-c", "virsh list | grep "+name).Run()
		if err == nil {
			status = "inactive"
		} else {
			status = "active"
		}
	} else if err.Error() == "exit status 1" {
		status = "NotFound"
	} else {
		fmt.Print("error: command exection error for checking vm status\n")
		os.Exit(1)
	}
	return status
}

// CheckFile is checking file exsist
func CheckFile(path string) (status string) {
	err := exec.Command("find", path).Run()

	if err != nil {
		status = "AlreadyExist"
	} else if err.Error() == "exit code 1" {
		status = "NotFound"
	} else {
		fmt.Print("error: command exection error for checking vm status\n")
		os.Exit(1)
	}
	return status
}
