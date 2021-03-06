package process

import ()

type YamlFile struct {
	Vm []VM `yaml:"vms"`
}

type VM struct {
	Name      string   `yaml:name`
	Count     int      `yaml:count`
	Vcpus     int      `yaml:vcpus`
	Memory    int      `yaml:memory`
	Disk      Disk     `yaml:disk`
	Graphics  string   `yaml:graphics`
	Location  string   `yaml:location`
	ExtraArgs string   `yaml:"extra-args"`
	Cdrom     string   `yaml:cdrom`
	Network   []string `yaml:network`
}

type Disk struct {
	Path       string `yaml:path`
	Format     string `yaml:format`
	Size       int    `yaml:size`
	ImportDisk bool   `yaml:import`
}
