package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// twonumberCmd represents the twonumber command
var twonumberCmd = &cobra.Command{
	Use:   "check [num1] [num2]",
	Short: "Add two numbers",
	Long:  "This command takes two numbers as arguments and prints their sum.",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		num1, err1 := strconv.Atoi(args[0])
		num2, err2 := strconv.Atoi(args[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Error: Both arguments must be integers.")
			return
		}
		fmt.Printf("The sum of %d and %d is %d\n", num1, num2, num1+num2)
	},
}

func init() {
	rootCmd.AddCommand(twonumberCmd)
}
