package process

import (
	"fmt"
	"os/exec"
)

func Build () {
	data := ReadYAML()
	_ = CheckYaml(data)
	BuildVM(data)
	fmt.Print("\n")
}

func BuildVM (data YamlFile) {
	for i := 0; i < len(data.Vm); i++ {
		arg := []string{
			"--name", data.Vm[i].Name,
			"--vcpus", data.Vm[i].Vcpus,
			"--memory", data.Vm[i].Memory,
		}

		if data.Vm[i].Disk.ImportDisk == true {
			importDisk := "path="+data.Vm[i].Disk.Path
			arg = append(arg, "--disk")
			arg = append(arg, importDisk)
			arg = append(arg, "import")
		} else {
			importDisk := "path="+data.Vm[i].Disk.Path+",format="+data.Vm[i].Disk.Format+",size="+data.Vm[i].Disk.Size
                        arg = append(arg, "--disk")
                        arg = append(arg, importDisk)
		}

		if data.Vm[i].Cdrom != "" {
			arg = append(arg, "--cdrom")
			arg = append(arg, data.Vm[i].Cdrom)
		} else {
			arg = append(arg, "--location")
			arg = append(arg, data.Vm[i].Location)
			if data.Vm[i].ExtraArgs != "" {
				arg = append(arg, "--extra-args")
				arg = append(arg, data.Vm[i].ExtraArgs)
			}
		}

		for j := 0; j < len(data.Vm); j++{
			arg = append(arg, "--network")
			arg = append(arg, data.Vm[i].Network[j])
		}

		cmd := exec.Command("virt-install", arg...)
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}
	return
}
