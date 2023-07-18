/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/Serenity0204/LVCS/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "staging the files",
	Long:  `add command will accept unlimited amount of files as the subcommand and will track them into staging.txt`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		dir += "\\.lvcs"
		if err != nil {
			fmt.Println("error retrieving .lvcs directory at:", dir)
			return
		}
		lvcsMan := internal.NewLVCSManager(dir)
		exists, err := lvcsMan.LVCSExists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if !exists {
			fmt.Println(".lvcs directory does not exist, add failed")
			return
		}
		msg, err := lvcsMan.Execute("add", args)
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
	rootCmd.AddCommand(addCmd)
}
