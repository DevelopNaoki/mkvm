package process

import (
	"os/exec"
	"strconv"
	"strings"
)

func Build() {
	data := ReadYAML()
	_ = CheckYaml(data)
	for i := 0; i < len(data.Vm); i++ {
		BuildVM(data.Vm[i])
	}
}

func BuildVM(data VM) {
	if data.Count == 0 {
		data.Count = 1
	}

	for i := 0; i < data.Count; i++ {
		arg := []string{"--name"}

		if data.Count == 1 {
			arg = append(arg, data.Name)
		} else {
			arg = append(arg, data.Name+strconv.Itoa(i))
		}
		arg = append(arg, "--vcpus")
		arg = append(arg, strconv.Itoa(data.Vcpus))
		arg = append(arg, "--memory")
		arg = append(arg, strconv.Itoa(data.Memory))

		if data.Disk.ImportDisk == true {
			importDisk := "path=" + data.Disk.Path
			arg = append(arg, "--disk")
			arg = append(arg, importDisk)
			arg = append(arg, "--import")
		} else {
			if data.Count > 1 {
				var disk [2]string
				slice := strings.Split(data.Disk.Path, ".")
				for i, str := range slice {
					disk[i] = str
				}
				createDisks := "path=" + disk[0] + strconv.Itoa(i) + "." + disk[1] + ",format=" + data.Disk.Format + ",size=" + strconv.Itoa(data.Disk.Size)
				arg = append(arg, "--disk")
				arg = append(arg, createDisks)
			} else {
				createDisk := "path=" + data.Disk.Path + ",format=" + data.Disk.Format + ",size=" + strconv.Itoa(data.Disk.Size)
				arg = append(arg, "--disk")
				arg = append(arg, createDisk)
			}
		}

		if data.Cdrom != "" {
			arg = append(arg, "--cdrom")
			arg = append(arg, data.Cdrom)
		} else {
			arg = append(arg, "--location")
			arg = append(arg, data.Location)
			if data.ExtraArgs != "" {
				arg = append(arg, "--extra-args")
				arg = append(arg, data.ExtraArgs)
			}
		}

		for j := 0; j < len(data.Network); j++ {
			arg = append(arg, "--network")
			arg = append(arg, data.Network[j])
		}

		cmd := exec.Command("virt-install", arg...)
		err := cmd.Start()
		if err != nil {
			panic(err)
		}
	}
	return
}
