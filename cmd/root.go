package cmd

import (
	"os"

	"github.com/rafamaciel/git-ignore/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-ignore",
	Short: "A simple git ignore manager",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		list, _ := cmd.Flags().GetBool("list")
		delete, _ := cmd.Flags().GetBool("delete")
		wd, _ := cmd.Flags().GetString("workdir")

		if wd == "" {
			wd, _ = os.Getwd()
		}

		if list {
			internal.ListIgnoredFiles(wd)
			return
		}

		if delete {
			if len(args) == 0 {
				return
			}
			for _, f := range args {
				if ok, _ := internal.IsFileInGitignore(wd, f); ok {
					_ = internal.RemovePatternFromGitignore(wd, f)
				}
			}
			return
		}

		if len(args) == 0 {
			return
		}

		for _, entry := range args {
			_, _ = internal.AddFileToGitignore(wd, entry)
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("list", "l", false, "List files on gitignore")
	rootCmd.Flags().BoolP("delete", "d", false, "Delete a file from gitignore")
	rootCmd.Flags().BoolP("match", "m", false, "Check if files match")
	rootCmd.Flags().String("workdir", "", "Working directory")
}
