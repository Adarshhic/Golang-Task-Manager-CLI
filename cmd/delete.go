
package cmd

import (
	"fmt"
	"strconv"

	"Project3_cli/Internal/Task"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [task-id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", args[0])
		}

		if err := task.DeleteTask(id); err != nil {
			return err
		}

		fmt.Printf("Task %d deleted successfully\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
