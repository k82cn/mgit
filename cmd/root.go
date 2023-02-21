/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mgit",
	Short: "The command lines to manage multiple repositories of solutions.",
	Long: `The command lines to manage multiple repositories of solutions. It will load the configuration from
${HOME}/.mgit by default, and the environment value ${MGIT_CONF} can also be used to set up the configuration.

    $ cat << EOF | tee ${HOME}/.mgit
    current-solution: openbce
    solutions:
      - name: openbce
        git_server: "git@github.com:"
        user: k82cn
        components:
          - name: device-manager
            git_path: openbce/device-manager
            module_path: openbce.io/device-manager
          - name: flame
            git_path: openbce/flame
            module_path: openbce.io/flame
          - name: kcache
            git_path: openbce/kcache
            module_path: openbce.io/kcache
`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
