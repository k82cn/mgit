/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package apis

const (
	ConfPathEnv     = "MGIT_CONF"
	DefaultConfName = ".mgit"
)

type Configuration struct {
	CurrentSolution string     `yaml:"current-solution,omitempty"`
	Workspace       string     `yaml:"workspace,omitempty"`
	Solutions       []Solution `yaml:"solutions,omitempty"`
}

type Solution struct {
	Name       string      `yaml:"name,omitempty"`
	Type       *string     `yaml:"type,omitempty"`
	GitServer  string      `yaml:"git_server,omitempty"`
	User       *string     `yaml:"user,omitempty"`
	Workspace  *string     `yaml:"workspace,omitempty"`
	Components []Component `yaml:"components,omitempty"`
}

type Component struct {
	Name         string  `yaml:"name,omitempty"`
	Type         *string `yaml:"type,omitempty"`
	User         *string `yaml:"user,omitempty"`
	GitPath      string  `yaml:"git_path,omitempty"`
	ModulePath   string  `yaml:"module_path,omitempty"`
	MainBranch   *string `yaml:"main_branch,omitempty"`
	BuildCommand *string `yaml:"build_command,omitempty"`
}
