/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "print current time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("now called")
		// date format
		cmdStruct := exec.Command("date", "+%Y-%m-%d %H:%M:%S")
		out, err := cmdStruct.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
