package model

type YamlFile struct {
	Hosts   []Host    `yaml:"hosts"`
	Vm      []VM      `yaml:"vms"`
	Disk    []Disk    `yaml:"disks"`
	Network []Network `yaml:"networks"`
}

type Host struct {
	Hostname          string `yaml:"hostname"`
	Port              string `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	IdentityKey       string `yaml:"identity-key"`
	PasswordAuthorize bool   `yaml:"password-authtorirze"`
	Hypervisor        string `yaml:"hypervisor"`
}

type VM struct {
	Name      string   `yaml:"name"`
	Host      string   `yaml:"host"`
	Count     int      `yaml:"count"`
	Vcpus     int      `yaml:"vcpus"`
	Memory    int      `yaml:"memory"`
	Disk      []string `yaml:"disk"`
	Graphics  string   `yaml:"graphics"`
	Location  string   `yaml:"location"`
	ExtraArgs string   `yaml:"extra-args"`
	Cdrom     string   `yaml:"cdrom"`
	Network   []string `yaml:"network"`
}

type Disk struct {
	Name       string `yaml:"name"`
	Host       string `yaml:"host"`
	Path       string `yaml:"path"`
	Format     string `yaml:"format"`
	Size       int    `yaml:"size"`
	ImportDisk bool   `yaml:"import"`
}

type Network struct {
	Name      string `yaml:"name"`
	Host      string `yaml:"host"`
	Interface string `yaml:"interface"`
	Type      string `yaml:"type"`
}
