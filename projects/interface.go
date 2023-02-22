/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package projects

import (
	"github.com/k82cn/mgit/apis"
	"sync"
)

type Command struct {
	Command   string
	Arguments []string
}

type Project interface {
	Name() string
	Dir(ws string) string
	PostUpdate() []Command
}

type NewProjectFunc func(component *apis.Component) Project

var pluginMap = make(map[string]NewProjectFunc)
var pluginMapMutex sync.Mutex

func New(c *apis.Component) Project {
	pluginMapMutex.Lock()
	defer pluginMapMutex.Unlock()

	if p, found := pluginMap[*c.Type]; found {
		return p(c)
	}

	return NewEmptyProject(c)
}

func register(name string, proj NewProjectFunc) {
	pluginMapMutex.Lock()
	defer pluginMapMutex.Unlock()

	pluginMap[name] = proj
}
