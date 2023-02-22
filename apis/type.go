/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package apis

const (
	ConfPathEnv     = "MGIT_CONF"
	DefaultConfName = ".mgit"
)

type Configuration struct {
	CurrentSolution string     `yaml:"current-solution"`
	Workspace       string     `yaml:"workspace"`
	Solutions       []Solution `yaml:"solutions"`
}

type Solution struct {
	Name       string      `yaml:"name"`
	Type       *string     `yaml:"type"`
	GitServer  string      `yaml:"git_server"`
	User       *string     `yaml:"user"`
	Workspace  *string     `yaml:"workspace"`
	Components []Component `yaml:"components"`
}

type Component struct {
	Name         string  `yaml:"name"`
	Type         *string `yaml:"type"`
	User         *string `yaml:"user"`
	GitPath      string  `yaml:"git_path"`
	ModulePath   string  `yaml:"module_path"`
	MainBranch   *string `yaml:"main_branch"`
	BuildCommand *string `yaml:"build_command"`
}
