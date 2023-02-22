/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package projects

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/k82cn/mgit/apis"
)

const GoProjectName = "go"

func init() {
	register(GoProjectName, NewGoProject)
}

func NewGoProject(c *apis.Component) Project {
	return &GoProject{component: c}
}

type GoProject struct {
	component *apis.Component
}

func (g *GoProject) Dir(ws string) string {
	if len(ws) == 0 {
		ws = os.Getenv("GOPATH")
	}

	return strings.Join([]string{ws, "src", g.component.ModulePath}, string(filepath.Separator))
}

func (g *GoProject) Name() string {
	return GoProjectName
}

func (g *GoProject) PostUpdate() []Command {
	return []Command{
		{
			Command:   "go",
			Arguments: []string{"mod", "tidy"},
		},
		{
			Command:   "go",
			Arguments: []string{"mod", "vendor"},
		},
	}
}
