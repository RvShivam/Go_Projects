package cmd

import (
	"fmt"
	"strconv"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// updatestatusCmd represents the updatestatus command
var updatestatusCmd = &cobra.Command{
	Use:   "updatestatus [id] [status]",
	Short: "Update Status of task",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("task id must be a number")
		}
		status := args[1]
		if (status == "todo") || (status == "done") || (status == "in-progress") {
			if err := internal.UpdateTaskStatus(id, status); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("invalid status. Status must be either \ttodo or \tdone")
		}
		fmt.Printf("Status of task %d updated to status: %s\n", id, status)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updatestatusCmd)
}
