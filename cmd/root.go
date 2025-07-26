package cmd

import (
	"fmt"
	"gitactivitytracker/activity"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitcheck",
	Short: "git check is a tool to track git user activity",
	Long:  `This is a git activity tracking CLI tool that helps to get a user activity for a repo`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("requires at least one arg")
		}
		if !isValidUser(args[0]) {
			return fmt.Errorf("invalid user")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		u := activity.UserActResponse{}
		activ, err := u.GetUserActivity(args[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Println(activ)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func isValidUser(user string) bool {
	if user == "" {
		return false
	}
	return true
}
