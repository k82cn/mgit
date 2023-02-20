/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download all related repositories of the solution",
	Long:  `Download all related repositories of the solution`,
	Run: func(cmd *cobra.Command, args []string) {
		sol, err := loadConfiguration()
		if err != nil {
			panic(err)
		}

		for _, repo := range sol.Components {
			cmd := exec.Command("git", "clone",
				strings.Join([]string{sol.GitServer, repo.GitPath}, "/"),
				"-o",
				strings.Join([]string{*sol.GoPath, repo.ModulePath}, "/"),
			)
			if err := cmd.Run(); err != nil {
				cmd.CombinedOutput()
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}
