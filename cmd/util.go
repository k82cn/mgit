/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	yamlv3 "gopkg.in/yaml.v3"

	"github.com/k82cn/mgit/apis"
	"github.com/k82cn/mgit/projects"
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

func loadConfiguration() (*apis.Configuration, error) {
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

	return res, nil
}

func loadSolution() (*apis.Solution, error) {
	conf, err := loadConfiguration()
	if err != nil {
		return nil, err
	}

	var sol *apis.Solution
	for _, s := range conf.Solutions {
		if conf.CurrentSolution == s.Name {
			sol = &s
			break
		}
	}

	if sol == nil {
		return nil, fmt.Errorf("current solution not found")
	}

	setDefault(sol, conf)

	return sol, nil
}

func setDefault(res *apis.Solution, conf *apis.Configuration) {
	if res.Workspace == nil {
		res.Workspace = &conf.Workspace
	}

	if res.User == nil {
		user := os.Getenv("USER")
		res.User = &user
	}

	if res.Type == nil {
		goproj := projects.GoProjectName
		res.Type = &goproj
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
		if res.Components[i].Type == nil {
			res.Components[i].Type = res.Type
		}
	}

}
