package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCMD is the base command when it is called without subcommands
var rootCmd = &cobra.Command{
	Use:   "expenseTracker",
	Short: "a simple, declarative, expense tracker",
	Long:  "This will save all your data to a CSV file on disk, and allow updating and modifying old entries as needed.",
	// Run defines what happens if you run the base command without args
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to the Expense Tracker!")
		fmt.Println("Run `expenseTracker --help` for available commands.")
	},
}

// Execute executes the root command
// when called by main.main()
func Execute() error {
	return rootCmd.Execute()
}
