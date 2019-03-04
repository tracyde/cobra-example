package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Example usage
// cobra-example command arg --flag

func parseFlag(f bool) string {
	if f {
		return "flag"
	}
	return "none"
}

func main() {

	// Flag create variable for command flag
	var Flag bool

	var rootCmd = &cobra.Command{
		Use:   "cobra-example",
		Short: "Example showing how to use Cobra",
		Long: `A brief example showing how to use Cobra created by spf13 and friends in Go.
		Complete documentation is available at https://github.com/spf13/cobra`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to cobra-example")
		},
	}

	var commandCmd = &cobra.Command{
		Use:   "command",
		Short: "Dummy command to show how to create a command",
		Long:  `All your base are belong to us`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to cobra-example")
			fmt.Println("you executed the command: command")
			fmt.Println("and passed in the arg: arg") // TODO dynamically print args
			fmt.Printf("flags: ")
			fmt.Printf("%v were passed\n", parseFlag(Flag))
		},
	}

	// Attach flag to command
	commandCmd.Flags().BoolVarP(&Flag, "flag", "f", false, "dummy flag")

	// Add command to the root command
	rootCmd.AddCommand(commandCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
