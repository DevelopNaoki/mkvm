package process

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// ReadYAML is read yaml and encode
func ReadYAML() (data YamlFile) {
	buf, err := ioutil.ReadFile("mkvm.yml")
	if err != nil {
		Error(30, "faild read file")
		return
	}

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}

	return
}

func CheckYaml(data YamlFile) (err int) {
	for i := 0; i < len(data.Vm); i++ {
	}
	return
}
