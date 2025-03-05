package executor

import (
	"fmt"

	"github.com/darkb0ts/BugFlowCandy/internal/config"
	"github.com/darkb0ts/BugFlowCandy/internal/utils"
	"github.com/darkb0ts/BugFlowCandy/model"
)

type BugRunner interface {
	Start() (bool, error)
	Run(cmd string) (string, error)
	Exit() (bool, error)
}

type Linear struct {
	Tools []config.Tool
}

func (l *Linear) Start() (bool, error) {
	if model.SharedData.Log {
		utils.LogInfo("Logging enabled")
	}
	if (model.SharedData.ConfigFile != "") && (model.SharedData.URL != "") && utils.CheckDirExists(model.SharedData.ConfigFile) {
		message, err := config.LoadConfig(model.SharedData.ConfigFile)
		if err != nil {
			fmt.Printf("Error loading config: %v\n", message)
			return false, nil
		}
		config.LoadData()

	}
	return true, nil
}

func (l *Linear) Run(cmd string) (string, error) {
	fmt.Printf("Running command: %s\n", cmd)
	return "", nil
}

func (l *Linear) Exit() (bool, error) {
	fmt.Println("Exiting Linear Runner")
	return true, nil
}

func NewLinearRunner(tools []config.Tool) *Linear {
	return &Linear{Tools: tools}
}
