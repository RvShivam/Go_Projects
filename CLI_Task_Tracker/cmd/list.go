package cmd

import (
	"fmt"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all existing tasks",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		tasks, err := internal.ListTasks()
		if err != nil {
			return err
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found")
			return nil
		}
		fmt.Printf("%-4s %-25s %-12s %-16s\n", "ID", "TITLE", "STATUS", "CREATED AT")
		for _, task := range tasks {
			fmt.Printf(
				"%-4d %-25s %-12s %-16s\n",
				task.ID,
				task.Title,
				task.Status,
				task.CreatedAt.Format("2006-01-02 15:04"),
			)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
