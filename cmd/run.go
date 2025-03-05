package cmd

import (
	"fmt"

	"github.com/darkb0ts/BugFlow/internal/utils"
	"github.com/spf13/cobra"
)

var (
	configFile string
	delayTime  int
	numThreads int
	outputDir  string
	url        string
	method     string // Linear, Tree, etc.
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute OSINT workflows from a YAML file",
	Run: func(cmd *cobra.Command, args []string) {
		// create is a function that creates a directory
		// err := utils.CreateDir(outputDir)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		commands := []string{
			"echo Hello",
			"rm -rf /",
			"ssh pi@192.168.1.1",
			"reboot",
			"ls -la",
		}
		for _, cmd := range commands {
			output, err := utils.ExecuteCommand(cmd)
			if err != nil {
				fmt.Printf("Command '%s' failed: %v\n", cmd, err)
			} else {
				fmt.Printf("Command '%s' output: %s\n", cmd, output)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&configFile, "file", "f", "", "YAML file containing workflows (required)")
	runCmd.Flags().StringVarP(&outputDir, "output", "o", "bugflowouput", "Output directory")
	runCmd.Flags().StringVarP(&url, "url", "u", "", "URL to scan target (required)")
	runCmd.Flags().IntVarP(&numThreads, "threads", "t", 2, "Number of goroutines")
	runCmd.Flags().IntVarP(&delayTime, "delay", "d", 3, "Delay time (ms) between executions")
	runCmd.Flags().StringVarP(&method, "method", "m", "Linear", "Default is Linear Method, if you want to use the scan method, use this flag[Linear,Tree]")
	runCmd.MarkFlagRequired("file")
	runCmd.MarkFlagRequired("url")

}
