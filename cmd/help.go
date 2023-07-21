package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// helpCmd represents the help command
var helpCmd = &cobra.Command{
	Use:   "help",
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
			content, err := os.ReadFile(dir + "/docs/list.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(content))
			return
		}
		if len(args) == 1 && args[0] == "detail" {
			content, err := os.ReadFile(dir + "/docs/detail.txt")
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(string(content))
			return
		}
		fmt.Println("unknown command")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)
}
