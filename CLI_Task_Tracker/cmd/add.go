package cmd

import (
	"fmt"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task name]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		title := args[0]

		if err := internal.AddTask(title); err != nil {
			return err
		}

		fmt.Println("Task added:", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
