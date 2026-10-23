
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"Project3_cli/Internal/Task"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		description := args[0]

		if err := task.AddTask(description); err != nil {
			return err
		}

		fmt.Printf("Task added successfully: %s\n", description)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
