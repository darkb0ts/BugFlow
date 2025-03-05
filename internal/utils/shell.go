package utils

import (
	"bytes"
	"errors"
	"os/exec"
	"strings"
)

func ExecuteCommand(command string) (string, error) {
	for _, forbidden := range blacklistedCommands {
		if strings.Contains(command, forbidden) {
			return "", errors.New("forbidden command found in the provided command")
		}
	}
	cmd := exec.Command("sh", "-c", command)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
