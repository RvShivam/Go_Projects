package cmd

import (
	"fmt"
	"strconv"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("task id must be a number")
		}

		if err := internal.DeleteTask(id); err != nil {
			return err
		}

		fmt.Println("Deleted task", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
