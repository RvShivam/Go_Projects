package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/RvShivam/Go_Projects/github-user-activity/internal"
	"github.com/spf13/cobra"
)

var (
	limit   int
	jsonOut bool
	showAll bool
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Fetch and display a user's GitHub activity",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		username, err := internal.LoadUser()
		if err != nil {
			return err
		}

		events, err := internal.GetActivity(username)
		if err != nil {
			return err
		}

		if showAll {
			lines := internal.GroupAndFormat(events)

			for _, line := range lines {
				fmt.Printf("- %s\n", line)
			}

		}
		if limit > 0 && len(events) > limit {
			events = events[:limit]
		}

		// JSON output
		if jsonOut {
			out, err := json.MarshalIndent(events, "", "  ")
			if err != nil {
				return err
			}
			fmt.Println(string(out))
			return nil
		}

		if len(events) == 0 {
			fmt.Println("No recent public activity found.")
			return nil
		}

		lines := internal.GroupAndFormat(events)

		for _, line := range lines {
			fmt.Printf("- %s\n", line)
		}

		return nil
	},
}

func init() {
	getCmd.Flags().IntVarP(
		&limit,
		"limit",
		"l",
		10,
		"Limit the number of activities displayed",
	)

	getCmd.Flags().BoolVar(
		&jsonOut,
		"json",
		false,
		"Output raw JSON instead of formatted text",
	)

	getCmd.Flags().BoolVar(
		&showAll,
		"all",
		false,
		"Show all event types (including unsupported ones)",
	)

	rootCmd.AddCommand(getCmd)
}
