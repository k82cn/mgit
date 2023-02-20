package cmd

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"

	"github.com/klausm/mgit/apis"
)

func loadConfiguration() (*apis.Solution, error) {
	confPath := os.Getenv(apis.ConfPathEnv)

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	res := &apis.Solution{}
	if err := yaml.Unmarshal(yamlFile, res); err != nil {
		return nil, err
	}

	if res.GoPath == nil {
		goPath := os.Getenv("GOPATH")
		res.GoPath = &goPath
	}

	if res.User == nil {
		user := os.Getenv("USER")
		res.User = &user
	}

	return res, nil
}
