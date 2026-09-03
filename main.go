package main

import (
	"fmt"

	"expenseTracker/cmd"
)

/*
Ok the plan is a simple tool to write expenses, those will be
then saved to a csv file with some buffering, and reloaded when
the program is started again.

I will use cobra from the start like 'argparse'

- add expense
- update past expenses
- delete past expenses
- view all past expenses
- view a summary of expenses
- view summary for a specific month (then we should also save a month
  when the expense is created)

After this is done I can add the interesting TUI with bubbletea
to make the package more interactive and with the eyecandy.
*/

type expense struct {
	description string
	category    string  // like shopping | food | loans
	value       float64 // user could pass int, in that case we convert it
	date        string  // automatically use current date?
}

func addExpense(e expense) {
	// open the expenses file and write directly?
	// or keep in memory until we close?
}

func main() {
	fmt.Println("hello! This is the start of the program!")
	cmd.Execute()
}
