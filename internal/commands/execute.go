package commands

import (
	"github.com/spf13/cobra"
	"polygun-cli/internal/commands/example"
)

// Execute runs the root command and handles any CLI execution exception. Additionally,
// all child command(s) are added to the root command.
func Execute(root *cobra.Command) {
	examples := &cobra.Group{ID: "examples", Title: "Example Commands"}

	root.AddGroup(examples)

	root.AddCommand(example.Command)

	if e := root.Execute(); e != nil {
		cobra.CheckErr(e)
	}
}
