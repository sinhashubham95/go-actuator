package core

import (
	"bytes"
	"github.com/sinhashubham95/go-actuator/commons"
	"runtime/pprof"
)

// GetThreadDump is used to get the dump of all the goroutines created
func GetThreadDump() ([]byte, error) {
	var buffer bytes.Buffer
	err := pprof.Lookup(commons.GoRoutines).WriteTo(&buffer, 1)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
