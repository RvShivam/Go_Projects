package cmd

import (
	"fmt"

	"github.com/RvShivam/Go_Projects/github-user-activity/internal"
	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user [account_name]",
	Short: "This  commands allows the user to add their profile.",
	Long: `This command can be used to add the github account whose activity is to viewed.
			USAGE:
					github-activity user [account_name]`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		accountName := args[0]

		existingUser, err := internal.LoadUser()
		if err != nil {
			return err
		}

		if err := internal.SaveUser(accountName); err != nil {
			return err
		}

		switch {
		case existingUser == "":
			fmt.Printf("User [%s] added successfully\n", accountName)

		case existingUser == accountName:
			fmt.Printf("User %s already exists.\n", existingUser)

		default:
			fmt.Printf("User %s changed to %s.\n", existingUser, accountName)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(userCmd)
}
