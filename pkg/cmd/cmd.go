package cmd

// uses a shared global cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func GetRootCmd() *cobra.Command {
	return rootCmd
}

// This is the type of function you want to register.
// It takes the command-line arguments and returns an error if something goes wrong.
type CommandFunc func() error

// RegisterCommand takes a root command, a name, and a function. It creates a new command
// whose name is the `name` argument and calls the provided function when run.
func RegisterCommand(name string, fn CommandFunc) {
	cmd := &cobra.Command{
		Use:   name,
		Short: fmt.Sprintf("Executes the %s function", name),
		RunE: func(cmd *cobra.Command, args []string) error {
			return fn()
		},
	}
	rootCmd.AddCommand(cmd)
}
