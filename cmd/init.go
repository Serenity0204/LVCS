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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init lvcs directory",
	Long:  `init lvcs directory to do other operations`,
	Run: func(cmd *cobra.Command, args []string) {
		dir, err := os.Getwd()
		dir += "\\.lvcs"
		if err != nil {
			fmt.Println("Error retrieving .lvcs directory at:", dir)
			return
		}
		if len(args) != 0 {
			fmt.Println("Too many arguments for init")
			return
		}
		lvcsMan := internal.NewLVCSManager(dir)
		exist, err := lvcsMan.LVCSExists()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if exist {
			fmt.Println(".lvcs directory already exists at:", dir)
			return
		}
		msg, err := lvcsMan.Execute("init", args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
