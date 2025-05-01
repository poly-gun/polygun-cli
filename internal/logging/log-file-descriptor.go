package logging

import (
	"os"
	"sync/atomic"
)

// LFD represents a "log-file-descriptor". The default is [os.Stdout].
var LFD atomic.Pointer[os.File]

func init() {
	LFD.Store(os.Stdout)
}
