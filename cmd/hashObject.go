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

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hashObject",
	Short: "A brief description of your command",
	Long:  ``,
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
			fmt.Println(".lvcs directory does not exist, hashObject failed")
			return
		}
		msg, err := lvcsMan.Execute("hashObject", args)
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
	rootCmd.AddCommand(hashObjectCmd)
}
