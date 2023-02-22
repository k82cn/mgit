/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package projects

import (
	"github.com/k82cn/mgit/apis"
	"path/filepath"
	"strings"
)

const EmptyProjectName = "empty"

func init() {
	register(EmptyProjectName, NewEmptyProject)
}

func NewEmptyProject(c *apis.Component) Project {
	return &EmptyProject{component: c}
}

type EmptyProject struct {
	component *apis.Component
}

func (g *EmptyProject) Dir(ws string) string {
	return strings.Join([]string{ws, "src", g.component.ModulePath}, string(filepath.Separator))
}

func (g *EmptyProject) Name() string {
	return EmptyProjectName
}

func (g *EmptyProject) PostUpdate() []Command {
	return nil
}
