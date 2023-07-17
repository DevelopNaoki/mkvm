package kvm

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/DevelopNaoki/mkvm/model"
)

// VMList executes virsh list [option]
func VMList(option model.VmListCmdOption) error {
	var args string

	if option.All {
		args = "--all"
	} else if option.Inactive {
		args = "--inactive"
	} else if option.Active {
		args = ""
	} else {
		return fmt.Errorf("option errro")
	}

	cmd := exec.Command("virsh", "list", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		return fmt.Errorf("command execution failed")
	}
	return nil
}

// VMOperation is doing VM operation: virsh [cont] [name]
func VMOperation(name string, cont string) {
	status, _ := CheckVMStatus(name)
	if status == "NotFound" {
		fmt.Print("error: " + name + " not found\n")
		return
	} else if status == "active" && cont == "start" {
		fmt.Print("error: " + name + " is already running\n")
		return
	} else if status == "inactive" && (cont == "shutdown" || cont == "destroy") {
		fmt.Print("error: " + name + " is already stopping\n")
		return
	} else {
		fmt.Print("error: Command execution failed\n")
		return
	}

	fmt.Print(cont + "VM \n")

	exec.Command("virsh", cont, name).Start()
	for {
		status, _ := CheckVMStatus(name)
		if (status == "active" && cont == "start") || (status == "inactive" && (cont == "shutdown" || cont == "destroy")) {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}
