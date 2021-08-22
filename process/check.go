package process

import (
	"fmt"
	"os"
	"os/exec"
)

// Check is the caller of each checking function
func Check() {
	var status string

	status = CheckVirtualization()
	fmt.Print("Check CPU virtualization...   " + status + "\n")

	status = CheckVirsh()
	fmt.Print("Check virsh command...        " + status + "\n")

	status = CheckVirtInstall()
	fmt.Print("Check virt-install command... " + status + "\n")
}

// CheckVirtualization is checking cpu virtualization enabled
func CheckVirtualization() (status string) {
	cmd := "grep -E 'svm|vmx' /proc/cpuinfo"
	err := exec.Command("sh", "-c", cmd).Run()

	if err == nil {
		status = "[\x1b[32mPass\x1b[0m]"
	} else {
		status = "[\x1b[31mFalse\x1b[0m]"
	}
	return
}

// CheckVirsh is checking virsh command installed
func CheckVirsh() (status string) {
	err := exec.Command("virsh", "version").Run()
	if err == nil {
		status = "[\x1b[32mPass\x1b[0m]"
	} else {
		status = "[\x1b[31mFalse\x1b[0m]"
	}
	return
}

// CheckVirtInstall is checking virt-install command installed
func CheckVirtInstall() (status string) {
	err := exec.Command("virt-install", "--version").Run()
	if err == nil {
		status = "[\x1b[32mPass\x1b[0m]"
	} else {
		status = "[\x1b[31mFalse\x1b[0m]"
	}
	return
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
	} else if err.Error() == "exit code 1"{
		status = "NotFound"
	} else {
		fmt.Print("error: command exection error for checking vm status\n")
                os.Exit(1)
	}
	return status
}
