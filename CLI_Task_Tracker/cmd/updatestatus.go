package cmd

import (
	"fmt"
	"strconv"

	"github.com/RvShivam/Go_Projects/CLI_Task_Tracker/internal"
	"github.com/spf13/cobra"
)

// updatestatusCmd represents the updatestatus command
var updatestatusCmd = &cobra.Command{
	Use:   "updatestatus",
	Short: "A brief description of your command",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("task id must be a number")
		}
		status := args[1]
		if (status == "todo") || (status == "done") {
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updatestatusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updatestatusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
