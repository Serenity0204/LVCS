package cmd

import (
	"github.com/spf13/cobra"
)

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Branching",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("branch", args)
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
