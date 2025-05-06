package level

import (
	"errors"
	"log/slog"
	"strings"

	"polygun-cli/internal/logging"

	"github.com/spf13/pflag"
)

// Type string that implements Cobra's Type interface for valid string enumeration values.
type Type string

const (
	Trace   Type = "trace"
	Debug   Type = "debug"
	Info    Type = "info"
	Notice  Type = "notice"
	Warning Type = "warning"
	Error   Type = "error"
)

// String is used both by fmt.Print and by Cobra in help text
func (o *Type) String() string {
	return string(*o)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (o *Type) Set(v string) error {
	switch strings.ToLower(v) {
	case "trace", "debug", "info", "notice", "warning", "error":
		*o = Type(v)

		return nil
	default:
		return errors.New("must be one of \"trace\", \"debug\", \"info\", \"notice\", \"warning\", \"error\"")
	}
}

// Type is only used in help text
func (o *Type) Type() string {
	return "[\"trace\"|\"debug\"|\"info\"|\"notice\"|\"warning\"|\"error\"]"
}

// Level - Exported constants representing [slog.Level].
//
// - Trace for tracing program's execution.
//
// - Debug for providing contextual information in debugging phase.
//
// - Info for informing about general system operations.
//
// - Notice for conditions that are not errors but might need handling.
//
// - Warning for warning conditions.
//
// - Error for error conditions.
//
// - Emergency for system-unusable conditions.
func (o *Type) Level() slog.Level {
	switch *o {
	case Trace:
		return logging.LevelTrace

	case Debug:
		return logging.LevelDebug

	case Info:
		return logging.LevelInfo

	case Notice:
		return logging.LevelNotice

	case Warning:
		return logging.LevelWarning

	case Error:
		return logging.LevelError

	default:
		return logging.LevelInfo
	}
}

// Runtime validation to ensure implementation satisfies the interface.
var _ pflag.Value = (*Type)(nil)
