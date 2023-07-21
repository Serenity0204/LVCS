package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "LVCS",
	Short: "Little Version Control System",
	Long:  `A Little Version Control System built in Golang Cobra with supporting init, add, commit, branch, hashObject, and catFile operations`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("type 'LVCS commands' to view all of the available commands or 'LVCS commands detail' to view the detail")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("list", "l", false, "list all of the avaible commands")
	rootCmd.Flags().BoolP("detail", "d", false, "list the detail of every command")
}
