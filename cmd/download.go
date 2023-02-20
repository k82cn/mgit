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

type downloadCmdOption struct {
	Force bool
}

var downloadCmdOpt downloadCmdOption

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download all related repositories of the solution",
	Long:  `Download all related repositories of the solution`,
	Run: func(cmd *cobra.Command, args []string) {
		sol, err := loadConfiguration()
		if err != nil {
			fmt.Printf("Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		for _, repo := range sol.Components {
			fmt.Printf("Start to download %s: ", repo.Name)

			src := strings.Join([]string{sol.GitServer, *sol.User, repo.Name}, "/")
			target := strings.Join([]string{*sol.GoPath, "src", repo.ModulePath}, string(filepath.Separator))
			if _, err := os.Stat(target); err == nil {
				if downloadCmdOpt.Force {
					if err := os.RemoveAll(target); err != nil {
						fmt.Println("Skipped.")
						continue
					}
				} else {
					fmt.Println("Skipped.")
					continue
				}
			}

			if err := os.MkdirAll(target, 0644); err != nil {
				fmt.Printf("Failed to create directories: %v\n", err)
				os.Exit(1)
			}

			cmd := exec.Command("git", "clone", src, target)
			if msg, err := cmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			cmd = exec.Command("git", "remote", "add", "upstream",
				strings.Join([]string{sol.GitServer, repo.GitPath}, "/"))
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
	rootCmd.AddCommand(downloadCmd)

	downloadCmd.Flags().BoolVarP(&downloadCmdOpt.Force, "force", "f", false, "Remove repository if exists")
}
