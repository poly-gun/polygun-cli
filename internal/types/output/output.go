package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
	"github.com/spf13/pflag"
)

// Type string that implements Cobra's [cobra.Type] interface for valid string enumeration values.
type Type string

const (
	JSON Type = "json"
	YAML Type = "yaml"
)

// String is used both by fmt.Print and by Cobra in help text
func (o *Type) String() string {
	return string(*o)
}

// Set must have pointer receiver so it doesn't change the value of a copy
func (o *Type) Set(v string) error {
	switch v {
	case "json", "yaml":
		*o = Type(v)
		return nil
	default:
		return errors.New("must be one of \"json\" or \"yaml\"")
	}
}

// Type is only used in help text
func (o *Type) Type() string {
	return "[\"json\"|\"yaml\"]"
}

// Runtime validation to ensure implementation satisfies the interface.
var _ pflag.Value = (*Type)(nil)

// Write serializes the provided datum into the specified format (JSON or YAML) and writes it to the given io.Writer.
// Returns an error if encoding fails or encounters an issue during writing.
func Write(writer io.Writer, format Type, pretty bool, datum interface{}) error {
	switch format {
	case JSON:
		encoder := json.NewEncoder(writer)
		if pretty {
			encoder.SetIndent("", "    ")
		}

		if e := encoder.Encode(datum); e != nil {
			exception := fmt.Errorf("failed to encode (json): %w", e)

			return exception
		}
	case YAML:
		if e := yaml.NewEncoder(writer, yaml.Indent(4)).Encode(datum); e != nil {
			exception := fmt.Errorf("failed to encode (yaml): %w", e)

			return exception
		}
	}

	return nil
}
