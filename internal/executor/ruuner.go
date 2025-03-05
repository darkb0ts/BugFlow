package runner

import (
	"fmt"

	"github.com/darkb0ts/BugFlowCandy/internal/config"
)

type Runner interface {
	Start() (bool, error)
	Run(cmd string) (string, error)
	Exit() (bool, error)
}

type Linear struct {
	Tools []config.Tool
}

func (l *Linear) Start() (bool, error) {
	fmt.Println("Starting Linear Runner")
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
