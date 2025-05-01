package logging

import "sync/atomic"

// Source is an atomic pointer that represents if source logging should be included. Usage should be
// evaluated as early as possible in the program's runtime.
var Source atomic.Pointer[bool]
