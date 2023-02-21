/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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

		fmt.Printf(" %-20s%-15s%-40s%-20s\n", "Name", "User", "GitServer", "Components#")
		for _, s := range conf.Solutions {
			setDefault(&s)
			p := " "
			if s.Name == conf.CurrentSolution {
				p = "*"
			}
			fmt.Printf("%s%-20s%-15s%-40s%-20d\n", p, s.Name, *s.User, s.GitServer, len(s.Components))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
