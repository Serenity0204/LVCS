package cmd

import (
	"github.com/spf13/cobra"
)

// dumpCmd represents the dump command
var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "Dump LVCS directory",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("dump", args)
	},
}

func init() {
	rootCmd.AddCommand(dumpCmd)
}
