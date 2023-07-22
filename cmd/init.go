package cmd

import (
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Init lvcs directory",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("init", args)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
