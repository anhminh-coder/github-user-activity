package cmd

import (
	"fmt"
	"github-activity/internal/activity"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "Fetch the recent activity of the specified Github user",
	Long: `Fetch the recent activity of the specified Github user by providing username

Example:
  github-activity <username>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return RunDisplayActivityCmd(args)
	},
	SilenceUsage:  true,
	SilenceErrors: true,
}

func RunDisplayActivityCmd(args []string) error {
	if len(args) == 0 {
		fmt.Println("please provide username")
		return nil
	}
	username := args[0]
	activities, err := activity.FetchGithubActivity(username)
	if err != nil {
		return err
	}
	activity.DisplayActivities(username, activities)
	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}
