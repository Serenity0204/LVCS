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

// branchCmd represents the branch command
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "lvcs branching",
	Long:  `lvcs branching that supports create, delete, checkout, check if exists, get current branch, get all exising branches functions`,
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
			fmt.Println(".lvcs directory does not exist, branching failed")
			return
		}
		msg, err := lvcsMan.Execute("branch", args)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(msg)
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
}
