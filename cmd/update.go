/*
Copyright © 2023 Klaus Ma <klaus@xflops.cn>
*/

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/k82cn/mgit/projects"
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update all related repositories in local",
	Long:  `Update all related repositories in local`,
	Run: func(cmd *cobra.Command, args []string) {
		sol, err := loadSolution()
		if err != nil {
			fmt.Printf("Failed to load configuration: %v\n", err)
			os.Exit(1)
		}

		for _, repo := range sol.Components {
			fmt.Printf("Start to update %s: ", repo.Name)
			project := projects.New(&repo)

			target := project.Dir(*sol.Workspace)
			checkoutCmd := exec.Command("git", "checkout", *repo.MainBranch)
			checkoutCmd.Dir = target
			if msg, err := checkoutCmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			fetchCmd := exec.Command("git", "fetch", "upstream")
			fetchCmd.Dir = target
			if msg, err := fetchCmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			mergeCmd := exec.Command("git", "merge",
				strings.Join([]string{"upstream", *repo.MainBranch}, "/"))
			mergeCmd.Dir = target
			if msg, err := mergeCmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			pushCmd := exec.Command("git", "push")
			pushCmd.Dir = target
			if msg, err := pushCmd.CombinedOutput(); err != nil {
				fmt.Println("Failed.")
				fmt.Println(string(msg))
				os.Exit(1)
			}

			for _, c := range project.PostUpdate() {
				postCmd := exec.Command(c.Command, c.Arguments...)
				postCmd.Dir = target
				if msg, err := postCmd.CombinedOutput(); err != nil {
					fmt.Println("Failed.")
					fmt.Println(string(msg))
					os.Exit(1)
				}
			}

			fmt.Println("Done.")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
