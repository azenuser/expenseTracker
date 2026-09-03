package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	expenseDescription string
	expenseValue       float64
)

// type expense struct {
// 	description string
// 	category    string  // like shopping | food | loans
// 	value       float64 // user could pass int, in that case we convert it
// 	date        string  // automatically use current date?
// }

// actually i would like a stryuct to this?
func addExpense(expense string) error {
	// io.write stuff
	fmt.Println("expense: ", expense)
	_, err := fmt.Println("expense saved to csv.")
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

// so to make this simpler, i can make a function with
// the signature that cobra is expecting.
func runAddExpense(cmd *cobra.Command, args []string) {
	// flag values here are populated from the cli inputs
	fmt.Println("Adding expense: ", expenseDescription, "with value: ", expenseValue)

	// but make a separte function for the logic
	err := saveToCsv(expenseDescription, expenseValue)
	if err != nil {
		fmt.Println("Error saving expense: ", err)
	}
}

func saveToCsv(desc string, value float64) error {
	// for these %s you need printf, and give the \n newline
	fmt.Printf("Saved to csv expense: %s (%.2f euro)\n", desc, value)
	return nil
}

// this is the actual command that will be used in the CLI
// like `appname add`
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an expense with the current date",
	// this can be 0 if you pass a flag?
	// Args:  cobra.ExactArgs(1),
	Args: cobra.NoArgs,
	// Run: func(cmd *cobra.Command, args []string) {
	// 	addExpense(expenseDescription)
	// },

	// is there an easier way to write the run aspect of the cmd?
	Run: runAddExpense,
}

func init() {
	// so here you need to pass the actual name
	// you used for the cobra command, not the func
	// it may call!
	rootCmd.AddCommand(addCmd)

	// and you also need to collect the data from user
	// and pass it to the application, this is the equivalentof the argparse
	// add argument line
	addCmd.Flags().StringVarP(&expenseDescription, "expense", "e", "generic expense", "add the text for the expense")
	// errorrs because this is a number and not a string?
	addCmd.Flags().Float64VarP(&expenseValue, "value", "v", 0.0, "The value of the expense.")
}
