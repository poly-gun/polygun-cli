package logging

import (
	"os"
	"path"
	"sync/atomic"
)

// Name is an atomic pointer to a string, typically used for thread-safe access or updates to a shared string value -- represents the
// name of the cli executable.
var Name atomic.Pointer[string]

// Executable returns the current value stored in the atomic string pointer Name. It is typically used for thread-safe access.
func Executable() string {
	return *(Name.Load())
}

func init() {
	v := path.Base(os.Args[0])

	Name.Store(&v)
}
