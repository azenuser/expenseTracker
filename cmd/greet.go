package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var nameFlag2 string

var greetCmd = &cobra.Command{
	Use:   "greet user",
	Short: "Prints greeting",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello, ", nameFlag2)
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	// StringVarP( pointer, longname, shortName, defaultkvalue, description)
	greetCmd.Flags().StringVarP(&nameFlag2, "name", "n", "unnamedPerson", "Name to greet")
}
