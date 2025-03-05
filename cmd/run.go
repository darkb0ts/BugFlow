package cmd

import (
	"fmt"

	"github.com/darkb0ts/BugFlowCandy/internal/config"
	"github.com/darkb0ts/BugFlowCandy/internal/executor"
	"github.com/darkb0ts/BugFlowCandy/model"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		tools := []config.Tool{} // Initialize tools as needed
		linearRunner := executor.NewLinearRunner(tools)
		test, err := linearRunner.Start()
		if err != nil {
			fmt.Println("Error starting executor:", err)
			return
		}
		fmt.Println("Executor started successfully:", test)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&model.SharedData.ConfigFile, "file", "f", "", "YAML file containing workflows (required)")
	runCmd.Flags().BoolVarP(&model.SharedData.Log, "log", "l", true, "Enable logging")
	runCmd.Flags().StringVarP(&model.SharedData.OutputDir, "output", "o", "bugflowouput", "Output directory")
	runCmd.Flags().StringVarP(&model.SharedData.URL, "url", "u", "", "URL to scan target (required)")
	runCmd.Flags().IntVarP(&model.SharedData.NumThreads, "threads", "t", 2, "Number of goroutines")
	runCmd.Flags().IntVarP(&model.SharedData.DelayTime, "delay", "d", 3, "Delay time (ms) between executions")
	runCmd.Flags().StringVarP(&model.SharedData.Method, "method", "m", "Linear", "Default is Linear Method, if you want to use the scan method, use this flag[Linear,Tree]")
	runCmd.MarkFlagRequired("file")
	runCmd.MarkFlagRequired("url")

}
