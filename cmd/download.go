/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"

	"github.com/k82cn/mgit/projects"
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
		solution, err := loadSolution()
		if err != nil {
			fmt.Printf("Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		for _, repo := range solution.Components {
			fmt.Printf("Start to download %s: ", repo.Name)
			project := projects.New(&repo)

			src := strings.Join([]string{solution.GitServer, *solution.User, repo.Name}, "/")
			target := project.Dir(*solution.Workspace)
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

			if err := os.MkdirAll(target, 0755); err != nil {
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
				strings.Join([]string{solution.GitServer, repo.GitPath}, "/"))
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
