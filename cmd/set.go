/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

type setOptions struct {
	CurrentSolution string
}

var setOpt setOptions

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Update configuration of mgit.",
	Long:  `Update configuration of mgit.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(setOpt.CurrentSolution) != 0 {
			conf, err := loadConfiguration()
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n.", err)
				os.Exit(1)
			}

			for _, s := range conf.Solutions {
				if s.Name == setOpt.CurrentSolution {
					conf.CurrentSolution = setOpt.CurrentSolution
				}
			}

			if conf.CurrentSolution != setOpt.CurrentSolution {
				fmt.Printf("Failed to find solution '%s' in configuration.\n", setOpt.CurrentSolution)
				os.Exit(1)
			}

			if err := saveConfiguration(conf); err != nil {
				fmt.Printf("Failed to save configuration: %v\n.", err)
				os.Exit(1)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&setOpt.CurrentSolution, "current-solution", "s", "", "Set current solution")
}
