package cmd

import (
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Log commit history",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("log", args)
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}
