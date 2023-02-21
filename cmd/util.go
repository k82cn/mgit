package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"

	"github.com/k82cn/mgit/apis"
)

func getConfPath() (string, error) {
	confPath := os.Getenv(apis.ConfPathEnv)
	if len(confPath) == 0 {
		confPath = strings.Join([]string{os.Getenv("HOME"), apis.DefaultConfName}, string(filepath.Separator))
	}

	if _, err := os.Stat(confPath); err != nil {
		return "", err
	}

	return confPath, nil
}

func loadConfiguration() (*apis.Solution, error) {
	confPath, err := getConfPath()
	if err != nil {
		return nil, err
	}

	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, err
	}

	res := &apis.Configuration{}
	if err := yamlv3.Unmarshal(yamlFile, res); err != nil {
		return nil, err
	}

	var sol *apis.Solution
	for _, s := range res.Solutions {
		if res.CurrentSolution == s.Name {
			sol = &s
			break
		}
	}

	if sol == nil {
		return nil, fmt.Errorf("current solution not found")
	}

	setDefault(sol)

	return sol, nil
}

func setDefault(res *apis.Solution) {
	if res.GoPath == nil {
		goPath := os.Getenv("GOPATH")
		res.GoPath = &goPath
	}

	if res.User == nil {
		user := os.Getenv("USER")
		res.User = &user
	}

	for i := range res.Components {
		mb := "main"
		if res.Components[i].MainBranch == nil {
			res.Components[i].MainBranch = &mb
		}
		if res.Components[i].User == nil {
			res.Components[i].User = res.User
		}
		if res.Components[i].BuildCommand == nil {
			buildCommand := "make"
			res.Components[i].BuildCommand = &buildCommand
		}
	}

}
