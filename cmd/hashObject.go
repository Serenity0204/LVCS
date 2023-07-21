package cmd

import (
	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hashObject",
	Short: "Store a file as an object by converting into OID",
	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("hashObject", args)
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
}
