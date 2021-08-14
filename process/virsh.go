package process

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

// VMList executes virsh list [option]
func VMList(option string) {
	var args []string

	if option == "active" {
		args = []string{
			"list",
		}
	} else if option == "all" || option == "inactive" {
		args = []string{
			"list",
			"--" + option,
		}
	} else {
		fmt.Print("error: Unknown option: "+ option +"\n")
		return
	}

	cmd := exec.Command("virsh", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()

	if err != nil {
		fmt.Print("error: Command execution failed\n")
	}
	return
}

// VMOperation is doing VM operation: virsh [cont] [name]
func VMOperation(name string, cont string) {
	status := VMSearch(name)
	if status == "NotFound" {
		fmt.Print("error: "+ name +" not found\n")
		return
	} else if status == "active" && cont == "start" {
		fmt.Print("error: "+ name +" is already running\n")
		return
	} else if status == "inactive" && (cont == "shutdown" || cont == "destroy") {
		fmt.Print("error: " + name + " is already stopping\n")
		return
	} else {
                fmt.Print("error: Command execution failed\n")
		return
	}

	fmt.Print(cont + "VM \n")

	arg := []string{
		cont,
		name,
	}

	exec.Command("virsh", arg...).Start()
	for {
		status := VMSearch(name)
		if (status == "active" && cont == "start") || (status == "inactive" && (cont == "shutdown" || cont == "destroy")) {
			fmt.Print("success\n")
			break
		}
		time.Sleep(time.Second)
		fmt.Print("*")
	}
	return
}

// VMSearch is serch VM by virsh command
func VMSearch(name string) (status string) {
	cmd := "virsh list --all | grep " + name
	err := exec.Command("sh", "-c", cmd).Run()

	if err != nil {
		status = "NotFound"
	} else {
		cmd = "virsh list | grep " + name
		err = exec.Command("sh", "-c", cmd).Run()
		if err != nil {
			status = "inactive"
		} else {
			status = "active"
		}
	}
	return status
}
