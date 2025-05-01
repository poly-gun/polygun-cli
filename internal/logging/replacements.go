package logging

import (
	"fmt"
	"log/slog"
	"strings"
)

func Replacements(groups []string, a slog.Attr) slog.Attr {
	switch a.Key {
	case slog.TimeKey:
		// Remove time from output
		return slog.Attr{}
	case slog.LevelKey:
		// Customize the name of the level key and the output string, including
		// custom level values.
		v := a.Value.Any().(slog.Level)

		// Renaming the log levels based on their priority
		switch {
		case v <= slog.Level(-8):
			a.Value = slog.StringValue("TRACE")
		case v <= slog.LevelDebug:
			a.Value = slog.StringValue("DEBUG")
		case v <= slog.LevelInfo:
			a.Value = slog.StringValue("INFO")
		case v <= slog.Level(2):
			a.Value = slog.StringValue("NOTICE")
		case v <= slog.LevelWarn:
			a.Value = slog.StringValue("WARNING")
		case v <= slog.LevelError:
			a.Value = slog.StringValue("ERROR")
		case v <= slog.Level(12):
			a.Value = slog.StringValue("EMERGENCY")
		default:
			a.Value = slog.StringValue("ERROR")
		}

	case slog.SourceKey:
		a.Key = "$"

		value := a.Value.String()[2 : len(a.Value.String())-1]
		partials := strings.Split(value, " ")
		value = strings.Join(partials[1:], ":")

		a.Value = slog.StringValue(fmt.Sprintf("file://%s", value))
	}

	return a
}
