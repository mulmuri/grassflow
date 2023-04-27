package execute

import (
	"errors"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type YamlConfig struct {
	TaskConfig []TaskConfig `yml:"grassflow"`
	WorkingDir string `yml:"working-dir"`
	TargetBranch string `yml:"target-branch"`
}

type TaskConfig struct {
	TaskName string `yml:"name"`
	Message  string `yml:"message"`
	Branch   string `yml:"branch"`
	Dir    []string `yml:"dir"`
	Match  []string `yml:"match"`
}


func ReadYML(fileName string, config *YamlConfig) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return err
	}

	if len(config.TaskConfig) == 0 {
		return errors.New("There is no task listed or file form is incorrect")
	}	

	return nil
}