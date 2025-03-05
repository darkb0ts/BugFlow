package cmd

import (
	"fmt"

	"github.com/darkb0ts/BugFlowCandy/internal/config"
	"github.com/darkb0ts/BugFlowCandy/internal/executor"
	"github.com/darkb0ts/BugFlowCandy/model"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the workflow with yaml file",
	Run: func(cmd *cobra.Command, args []string) {
		tools := []config.Tool{} // Initialize tools as needed
		linearRunner := executor.NewLinearRunner(tools)
		sss, err := linearRunner.Start()
		if err != nil {
			fmt.Println("Error starting executor:", err)
			return
		}
		fmt.Println("Executor started successfully:", sss)
	},
}

// Show the workflow

func init() {
	showCmd.Flags().StringVarP(&model.SharedData.URL, "url", "u", "", "URL to scan target (required)")
	showCmd.Flags().StringVarP(&model.SharedData.ConfigFile, "file", "f", "", "YAML file containing workflows (required)")
	rootCmd.AddCommand(showCmd)
}
