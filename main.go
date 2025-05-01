package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"strings"

	"github.com/spf13/cobra"
	"polygun-cli/internal/commands"
	"polygun-cli/internal/logging"
	"polygun-cli/internal/types/level"
)

var (
	version string = "0.0.0"  // See go linking for compile-time variable overwrites.
	commit  string = "n/a"    // See go linking for compile-time variable overwrites.
	date    string = "latest" // See go linking for compile-time variable overwrites.
)

// lvl represents the log-level flag set by a persisted global flag.
var lvl level.Type = "info"

// src represents the include-source-logging flag set by a persisted global flag.
var src bool = false

// // stderr represents the stderr flag set by persisted global flag.
// var stderr bool = false

func main() {
	// The PersistentPreRun and PreRun functions will be executed before Run. PersistentPostRun and PostRun will be executed
	// after Run. The Persistent*Run functions will be inherited by children if they do not declare their own. The *PreRun
	// and *PostRun functions will only be executed if the Run function of the current command has been declared. These
	// functions are run in the following order:
	//
	// - PersistentPreRun
	// - PreRun
	// - Run
	// - PostRun
	// - PersistentPostRun
	//
	// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md
	//

	var root = &cobra.Command{
		Use:         fmt.Sprintf("%s", logging.Executable()),
		Short:       "A cli tool [...]",
		Long:        "A cli tool [...]",
		Example:     "",
		Annotations: nil,
		Version:     version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			logging.Level.Set(lvl.Level())

			// Set the logging-file-descriptor if stderr.
			// if stderr {
			// 	logging.LFD.Store(os.Stderr)
			// }

			// Setup slog-specific logging.
			writer := logging.LFD.Load()
			options := &slog.HandlerOptions{AddSource: src, Level: &logging.Level, ReplaceAttr: logging.Replacements}
			handler := slog.NewJSONHandler(writer, options)

			logger := slog.New(handler)

			slog.SetDefault(logger)

			slog.Log(ctx, logging.LevelTrace, "Root", slog.Group("command",
				slog.String("name", cmd.Name()),
				slog.String("version", version),
				slog.String("commit", commit),
				slog.String("date", date),
			))

			return nil
		},
		// @todo Logic to check if a newer version is available
		// PreRun: func(cmd *cobra.Command, args []string) {},
		// Run: func(cmd *cobra.Command, args []string) {
		// 	if len(args) == 0 {
		// 		if e := cmd.Help(); e != nil {
		// 			panic(e)
		// 		}
		// 	}
		// },
		PostRun:           nil,
		CompletionOptions: cobra.CompletionOptions{},
		TraverseChildren:  true,
		Hidden:            false,
		SilenceErrors:     false,
		SilenceUsage:      false,
	}

	root.PersistentFlags().VarP(&lvl, "log-level", "z", "verbosity")
	root.PersistentFlags().BoolVarP(&src, "include-source-logging", "x", false, "display runtime caller locations")
	// root.PersistentFlags().BoolVarP(&stderr, "use-standard-error", "z", false, "redirect logs to stderr")

	commands.Execute(root)
}

func init() {
	version = strings.TrimSpace(version)
}
