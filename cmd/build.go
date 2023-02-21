/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the components accordingly",
	Long:  `Build the components accordingly`,
	Run: func(cmd *cobra.Command, args []string) {
		solution, err := loadSolution()
		if err != nil {
			fmt.Printf("Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		for _, repo := range solution.Components {
			fmt.Printf("Start to build %s: ", repo.Name)

			target := strings.Join([]string{*solution.GoPath, "src", repo.ModulePath}, string(filepath.Separator))

			cmd := exec.Command("/bin/bash", "-c", *repo.BuildCommand)
			cmd.Dir = target
			if msg, err := cmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			fmt.Println("Done.")
		}
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
