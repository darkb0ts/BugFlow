package cmd

import (
	"fmt"

	"github.com/darkb0ts/BugFlowCandy/internal/config"
	"github.com/darkb0ts/BugFlowCandy/internal/utils"
	"github.com/darkb0ts/BugFlowCandy/model"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the workflow",
	Run: func(cmd *cobra.Command, args []string) {
		if model.SharedData.Log {
			utils.LogInfo("Logging enabled")
		}
		if (model.SharedData.ConfigFile != "") && (model.SharedData.URL != "") && utils.CheckDirExists(model.SharedData.ConfigFile) {
			message, err := config.LoadConfig(model.SharedData.ConfigFile)
			if err != nil {
				fmt.Printf("Error loading config: %v\n", message)
				return
			}

		}
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
