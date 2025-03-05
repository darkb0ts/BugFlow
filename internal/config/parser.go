package config

import (
	"fmt"

	logger "github.com/darkb0ts/BugFlowCandy/internal/utils"
	viper "github.com/spf13/viper"
)

type Tool struct {
	Name    string `mapstructure:"name"`
	Command string `mapstructure:"command"`
	Output  string `mapstructure:"output"`
	Next    []Tool `mapstructure:"next,omitempty"`
}

type Linear struct {
	Tools []Tool `mapstructure:"tools"`
}

type Tree struct {
	Tools []Tool `mapstructure:"tools"`
}

type ToolsTree struct {
	Tools []Tool `mapstructure:"tools"`
}

func LoadConfig(filepath string) (bool, error) {
	viper.SetConfigName(filepath) // Just the name without extension
	viper.SetConfigType("yaml")   // Set the type explicitly
	viper.AddConfigPath(".")      // Add the directory where the file is stored

	if err := viper.ReadInConfig(); err != nil {
		logger.LogError("Error reading config file, %s", err)
	}
	return true, nil
}

func LoadData() (*[]Tool, error) {
	var toolsTree ToolsTree
	if err := viper.Unmarshal(&toolsTree); err != nil {
		logger.LogError("Unable to decode into struct, %v", err)
	}

	if len(toolsTree.Tools) == 0 {
		return nil, fmt.Errorf("no tools found in config")
	}

	for i, tool := range toolsTree.Tools {
		printToolDetails(i, tool) // Call printToolDetails for each tool
	}

	return &toolsTree.Tools, nil
}

func printToolDetails(i int, k Tool) {
	// Print the current tool's details
	fmt.Printf("%d %s %s %s\n", i, k.Command, k.Name, k.Output)

	// Iterate over the "next" tools and print their details recursively
	for _, nextTool := range k.Next {
		printToolDetails(i, nextTool) // Recursively print the nested tools
	}
}
