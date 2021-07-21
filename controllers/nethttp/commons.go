package nethttp

import (
	"github.com/sinhashubham95/go-actuator/commons"
	"github.com/sinhashubham95/go-actuator/core"
)

// EncodeJSON is used to encode the given interface into json bytes
var EncodeJSON = commons.EncodeJSON

// GetThreadDump is used to get the thread dump for the current runtime
var GetThreadDump = core.GetThreadDump
