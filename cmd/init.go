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
	Short: "Init lvcs directory",
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
	rootCmd.AddCommand(initCmd)
}
