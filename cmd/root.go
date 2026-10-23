/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"
    "Project3_cli/Internal/Task"
	"github.com/spf13/cobra"
)





var rootCmd = &cobra.Command{
    Use:   "Project3_cli",
    Short: "A CLI task manager",

    PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
        return task.LoadTasks()
    },
}



func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}



func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


