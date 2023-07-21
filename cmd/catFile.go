package cmd

import (
	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "catFile",
	Short: "Convert OID to file content",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("catFile", args)
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)
}
