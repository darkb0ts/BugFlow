package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "BugFlow",
	Long: `BugFlow is a command-line tool designed to help you run multiple bug bounty tools in a single platform.`,
}

func Execute() bool {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
	return true
}
