package process

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// VMList executes virsh list [option]
func VMList(option string) {
	var args string

	if option == "active" {
		args = "list"
	} else if option == "all" || option == "inactive" {
		args = "list --" + option
	} else {
		fmt.Print("error: Unknown option: " + option + "\n")
		return
	}

	cmd := exec.Command("virsh", args)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Print("error: Command execution failed\n")
		panic(err)
	}
	return
}

// VMOperation is doing VM operation: virsh [cont] [name]
func VMOperation(name string, cont string) {
	status := CheckVMStatus(name)
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
		status := CheckVMStatus(name)
		if (status == "active" && cont == "start") || (status == "inactive" && (cont == "shutdown" || cont == "destroy")) {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}
