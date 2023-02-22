/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the build command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all solutions in the configuration.",
	Long:  `List all solutions in the configuration.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := loadConfiguration()
		if err != nil {
			fmt.Printf("Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("  %-15s%-15s%-60s%-20s\n", "Name", "User", "GitServer", "Components#")
		for _, s := range conf.Solutions {
			setDefault(&s, conf)
			p := " "
			if s.Name == conf.CurrentSolution {
				p = "*"
			}
			fmt.Printf("%s %-15s%-15s%-60s%-20d\n", p, s.Name, *s.User, s.GitServer, len(s.Components))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
