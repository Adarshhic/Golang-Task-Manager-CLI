
package cmd

import (
	"fmt"
	"strconv"

	"Project3_cli/Internal/Task"
	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete [task-id]",
	Short: "Mark a task as completed",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid task ID: %s", args[0])
		}
		if err := task.CompleteTask(id); err != nil {
			return err
		}

		fmt.Printf("Task %d marked as completed\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
