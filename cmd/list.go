
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"Project3_cli/Internal/Task"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks := task.ListTasks()

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return nil
		}

		for _, t := range tasks {
			fmt.Printf("%d. %s | Completed: %t\n",
				t.ID, t.Description, t.Completed)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
