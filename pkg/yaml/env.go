package yaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Envs env map which hold values
type Envs struct {
	Env map[string]string
}

// FetchEnvVarsFromYaml reads env from file
func FetchEnvVarsFromYaml(filePath string) (newEnvs Envs, err error) {
	if filePath == "" {
		filePath = "env.yml"
	}

	data, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		return newEnvs, readErr
	}

	newEnvs = Envs{}

	err = yaml.Unmarshal([]byte(data), &newEnvs)
	if err != nil {
		log.Printf("error: %v", err)
		return newEnvs, err
	}

	return newEnvs, nil
}

// SetEnvVars set values to environment
func SetEnvVars(newEnvs Envs) {
	for k, v := range newEnvs.Env {
		fmt.Printf("Setting Env Var: %s\n", k)
		os.Setenv(k, v)
	}
}
