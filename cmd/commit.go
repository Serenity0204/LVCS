package cmd

import (
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit a version",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("commit", args)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
