package logging

import (
	"os"
	"sync/atomic"
)

// LFD represents a "log-file-descriptor" - where logs are written to.
var LFD atomic.Pointer[os.File]

func init() {
	LFD.Store(os.Stderr)
}
