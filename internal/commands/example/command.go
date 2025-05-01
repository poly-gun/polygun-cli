package example

import (
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"polygun-cli/internal/logging"
	"polygun-cli/internal/types/output"
)

// Runtime, command-specific flag(s).
var (
	name   = ""
	format = output.JSON
	pretty = false
)

var Command = &cobra.Command{
	Use:        "example",
	Aliases:    []string{},
	SuggestFor: nil,
	Short:      "The example's command short-description",
	Long:       "The example's command long-description -- value should be in full sentences, and can span multiple lines.",
	Example: strings.Join([]string{
		fmt.Sprintf("  %s", "# General command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\"", logging.Executable())),
		"",
		fmt.Sprintf("  %s", "# Extended usage demonstrating configuration of default(s)"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --name \"test-value\" --output json", logging.Executable())),
		"",
		fmt.Sprintf("  %s", "# Display help information and command usage"),
		fmt.Sprintf("  %s", fmt.Sprintf("%s example --help", logging.Executable())),
	}, "\n"),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		slog.DebugContext(ctx, "Example Log Message", slog.Group("command",
			slog.String("name", cmd.Name()),
			slog.Group("flags",
				slog.String("name", name),
				slog.String("output", format.String()),
			),
		))

		var datum = map[string]string{
			"name":   name,
			"output": format.String(),
		}

		var buffer bytes.Buffer
		if e := output.Write(&buffer, format, pretty, datum); e != nil {
			return e
		}

		fmt.Fprintf(os.Stdout, "%s", buffer.String())

		return nil
	},
	TraverseChildren: true,
}

func init() {
	flags := Command.Flags()

	flags.StringVarP(&name, "name", "n", "", "a required example named-string-flag")
	flags.BoolVarP(&pretty, "pretty", "p", false, "format output in a more human-readable format")
	flags.VarP(&format, "output", "o", "structured data format")
	if e := Command.MarkFlagRequired("name"); e != nil {
		if exception := Command.Help(); exception != nil {
			panic(exception)
		}
	}
}
