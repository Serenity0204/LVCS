package cmd

import (
	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var stageCmd = &cobra.Command{
	Use:   "stage",
	Short: "Staging",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("stage", args)
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)
}
