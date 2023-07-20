package cmd

import (
	"fmt"
	"os"

	"github.com/Serenity0204/LVCS/internal"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		dir += "\\.lvcs"
		if err != nil {
			fmt.Println("Error retrieving .lvcs directory at:", dir)
			return
		}
		lvcsMan := internal.NewLVCSManager(dir)
		exists, err := lvcsMan.LVCSExists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if !exists {
			fmt.Println(".lvcs directory does not exist, commit failed")
			return
		}
		msg, err := lvcsMan.Execute("commit", args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		art, err := lvcsMan.GetRandomASCIIArt()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(art + "\n\n")
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
