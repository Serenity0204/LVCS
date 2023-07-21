package cmd

import (
	"fmt"
	"os"

	"github.com/Serenity0204/LVCS/resources"
	"github.com/spf13/cobra"
)

// commandsCmd represents the help command
var commandsCmd = &cobra.Command{
	Use:   "commands",
	Short: "print out commands",
	Long:  `print out all of the availible commands with or without details`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = os.Stat(dir)
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) == 0 {
			fmt.Println(resources.LIST)
			return
		}
		if len(args) == 1 && args[0] == "detail" {
			fmt.Println(resources.DETAIL)
			return
		}
		fmt.Println("unknown command")
	},
}

func init() {
	rootCmd.AddCommand(commandsCmd)
}
