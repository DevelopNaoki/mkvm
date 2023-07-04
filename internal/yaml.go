package internal

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/DevelopNaoki/mkvm/model"
)

// ReadYAML is read yaml file and encode
func ReadYAML(filename string) (data model.YamlFile, _ error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return data, fmt.Errorf("faild read %s", filename)
	}

	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return data, fmt.Errorf("faild parse %s", filename)
	}

	return data, nil
}
