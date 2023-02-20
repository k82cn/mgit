package apis

const (
	ConfPathEnv = "MGIT_CONF"
)

type Solution struct {
	Name       string      `yaml:"name"`
	GitServer  string      `yaml:"git_server"`
	User       *string     `yaml:"user"`
	GoPath     *string     `yaml:"go_path"`
	Components []Component `yaml:"components"`
}

type Component struct {
	Name         string  `yaml:"name"`
	User         *string `yaml:"user"`
	GitPath      string  `yaml:"git_path"`
	ModulePath   string  `yaml:"module_path"`
	MainBranch   string  `yaml:"main_branch"`
	BuildCommand *string `yaml:"build_command"`
}
