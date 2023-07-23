/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "",

	Run: func(cmd *cobra.Command, args []string) {
		executeLVCSCommandHelper("restore", args)
	},
}

func init() {
	rootCmd.AddCommand(restoreCmd)
}
