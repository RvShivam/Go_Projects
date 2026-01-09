package cmd

import (
	"fmt"
	"strconv"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [id] [new task title]",
	Short: "Update an existing task",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("task id must be a number")
		}

		title := args[1]

		if err := internal.UpdateTask(id, title); err != nil {
			return err
		}

		fmt.Println("Updated task", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
