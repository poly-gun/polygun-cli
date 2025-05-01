package logging

import (
	"log/slog"
)

// Level is a dynamic [slog.LevelVar] variable that allows a Handler level to change dynamically. It implements [slog.Leveler] as well as a [slog.LevelVar.Set]
// method, and it is safe for use by multiple goroutines. The zero [slog.LevelVar] corresponds to [slog.LevelInfo].
var Level slog.LevelVar

const (
	LevelTrace     = slog.Level(-8)
	LevelDebug     = slog.LevelDebug
	LevelInfo      = slog.LevelInfo
	LevelNotice    = slog.Level(2)
	LevelWarning   = slog.LevelWarn
	LevelError     = slog.LevelError
	LevelEmergency = slog.Level(12)
)
